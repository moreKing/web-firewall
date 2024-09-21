package pty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility/dir"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var shellMode = "sh"

func init() {
	shell := g.Cfg().MustGet(context.Background(), "server.shellMode", "sh").String()
	if shell == "bash" {
		shellMode = "bash"
	}
}

// Shell 定义一个结构体 方便保存各种连接信息
type Shell struct {
	Websocket *websocket.Conn
	Pty       *os.File // linux一切皆文件 shell也是一个可读写的文件
	LogShell  *entity.LogShell
	FilePath  string
	Ctx       context.Context
	Cancel    context.CancelFunc

	ZModemSZ, ZModemRZ, ZModemSZOO bool
}

type WebMsg struct {
	Type string `json:"type"`
	Cols int    `json:"Cols"`
	Rows int    `json:"rows"`
	Data []byte `json:"data"`
}

func NewShell(websocket *websocket.Conn, ctx1 context.Context) (shell *Shell, err error) {
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		if err != nil {
			cancel()
		}
	}()
	// Create arbitrary command.
	c := exec.Command(shellMode)

	// Start the command with a pty.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		g.Log().Error(ctx1, "Failed to get user home directory", err)
		return nil, err
	}

	err = os.Chdir(homeDir)
	if err != nil {
		g.Log().Error(ctx1, "Failed to Chdir user home directory", err)
		return nil, err
	}

	ptmx, err := Start(c)
	if err != nil {
		g.Log().Error(ctx1, "Shell Start err:", err)
		return nil, err
	}

	defer func(ptmx *os.File) {
		if err != nil {
			_ = ptmx.Close()
		}
	}(ptmx)
	users, err := service.User().GetUserByID(ctx1, service.BizCtx().Get(ctx1).User.ID)
	if err != nil {
		return nil, err
	}

	nowTime := time.Now()
	logPath := g.Cfg().MustGet(ctx1, "server.shellLog", "/var/log/web-firewall").String()

	filepath := path.Join(logPath, fmt.Sprintf("%d/%d/%d", nowTime.Year(), nowTime.Month(), nowTime.Day()))
	//判断文件路径是否存在
	if !dir.DirExists(filepath) {
		_ = dir.CreateDir(filepath)
	}

	var logShell = &entity.LogShell{
		Loginname: users.Loginname,
		Username:  users.Username,
		ClientIp:  ghttp.RequestFromCtx(ctx1).GetClientIp(),
		UserId:    users.Id,
		Success:   true,
		Online:    true,

		Filename:  gonanoid.Must(),
		CreatedAt: time.Now().Unix(),
	}

	id, err := dao.LogShell.Ctx(ctx1).InsertAndGetId(do.LogShell{
		Loginname: logShell.Loginname,
		Username:  logShell.Username,
		ClientIp:  logShell.ClientIp,
		UserId:    logShell.UserId,
		Success:   logShell.Success,
		Online:    logShell.Online,
		Filename:  logShell.Filename,
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		return nil, err
	}
	logShell.Id = id
	return &Shell{
		Websocket:  websocket,
		LogShell:   logShell,
		Ctx:        ctx,
		Pty:        ptmx,
		Cancel:     cancel,
		FilePath:   filepath,
		ZModemSZ:   false,
		ZModemRZ:   false,
		ZModemSZOO: false,
	}, nil
}

// 启动相互转发的连接
func (s *Shell) Start() {
	go s.Send2Web()
	go s.Send2Shell()
	go s.Stop()

}
func (s *Shell) Send2Shell() {

	defer g.Log().Info(s.Ctx, "关闭连接")
	for {
		select {
		case <-s.Ctx.Done():
			return
		default:
			//read websocket msg  需要通过msgType 判断是传输类型
			dataType, wsData, err := s.Websocket.ReadMessage()
			if err != nil {
				g.Log().Error(s.Ctx, err)
				s.Cancel()
				return
			}
			//二进制数据 直接发送到pty通道
			if dataType == websocket.BinaryMessage {
				_, _ = s.Pty.Write(wsData)
				continue
			}

			msg := &WebMsg{}
			err = json.Unmarshal(wsData, msg)

			if err != nil {

				g.Log().Info(s.Ctx, "解析数据转化成json失败", err)
				msg.Type = "cmd" // return
			}
			switch msg.Type {
			case "resize":
				//err = m.Session.WindowChange(msg.Rows, msg.Cols)
				//if err != nil {
				//	fmt.Println("重设宽高失败：", err)
				//}
				err := Setsize(s.Pty, &Winsize{
					Rows: uint16(msg.Rows),
					Cols: uint16(msg.Cols),
				})
				if err != nil {
					fmt.Println("重设宽高失败：", err)
					return
				}
			case "stdin":
				_, err = s.Pty.Write(msg.Data)
				if err != nil {
					fmt.Println("发送pty数据失败：", err)
				}

			case "cmd":
				_, err = s.Pty.Write(wsData)
				if err != nil {
					fmt.Println("cmd pty发送数据失败：", err)
				}

			case "ignore":
				// 其他
				fmt.Println(string(msg.Data))
			}

		}
	}
}
func (s *Shell) Send2Web() {

	// 创建读写文件 用于审计数据存储
	fileTimeLine, err := os.OpenFile(path.Join(s.FilePath, s.LogShell.Filename+".tl"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		g.Log().Error(s.Ctx, "创建fileTimeLine文件失败")
		return
	}
	fileOperator, err := os.OpenFile(path.Join(s.FilePath, s.LogShell.Filename+".data"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		g.Log().Error(s.Ctx, "创建fileOperator文件失败")
		return
	}

	defer g.Log().Info(s.Ctx, "退出websocket shell")
	for {
		select {
		case <-s.Ctx.Done():
			return

		default:
			// n := mySSh.Stdout.buffer.Len()
			sshOut := make([]byte, 8192)
			n, err := s.Pty.Read(sshOut)
			if err != nil {
				s.Cancel()
				g.Log().Error(s.Ctx, err)
				return
			}

			//fmt.Println(string(sshOut[:n]))

			gbkOut := sshOut[:n]

			if n > 0 {
				//拿到 sshOut输出后，首先判断当前文件传输状态是否结束
				if s.ZModemSZOO {
					// 经过测试 centos7-8 使用的 lrzsz-0.12.20 在 sz 结束时会发送 ZModemSZEndOO
					// 而 deepin20 等自带更新的 lrzsz-0.12.21rc 在 sz 结束时不会发送 ZModemSZEndOO， 而前端 zmodemjs
					// 库只有接收到 ZModemSZEndOO 才会认为 sz 结束，固这里需判断 sz 结束时是否发送了 ZModemSZEndOO，
					// 如果没有则手动发送一个，以便保证前端 zmodemjs 库正常运行（如果不发送，会导致使用 sz 命令时无法连续
					// 下载多个文件）。
					s.ZModemSZOO = false
					if n < 2 {
						// 手动发送 ZModemSZEndOO
						_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZEndOO)
						_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: sshOut[:n]})
						// mySSh.Websocket.WriteJSON(&WebMsg{Type: "stdin", ReslutData: string(sshOut[:n])})
						fmt.Println("ssh-->web" + string(sshOut[:n]))
					} else if n == 2 {
						if sshOut[0] == ZModemSZEndOO[0] && sshOut[1] == ZModemSZEndOO[1] {
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZEndOO)
						} else {
							// 手动发送 ZModemSZEndOO
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZEndOO)
							_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: sshOut[:n]})
							fmt.Println("ssh-->web" + string(sshOut[:n]))
							// mySSh.Websocket.WriteJSON(&WebMsg{Type: "stdin", ReslutData: string(sshOut[:n])})

						}
					} else {
						if sshOut[0] == ZModemSZEndOO[0] && sshOut[1] == ZModemSZEndOO[1] {
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, sshOut[:2])
							_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: sshOut[2:n]})
							// mySSh.Websocket.WriteJSON(&WebMsg{Type: "stdin", ReslutData: string(sshOut[:n])})

						} else {
							// 手动发送 ZModemSZEndOO
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZEndOO)
							_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: sshOut[:n]})
							// mySSh.Websocket.WriteJSON(&WebMsg{Type: "stdin", ReslutData: string(sshOut[:n])})

						}
					}

				} else {
					//判断当前状态是否是文件上传或下载， 再判断包有没有文件传输的信息
					if s.ZModemSZ {
						//判断是否有结束标记
						if x, ok := IsContain(sshOut[:n], ZModemSZEnd); ok { //判断是完成下载
							s.ZModemSZ = false
							s.ZModemSZOO = true
							fmt.Println("结束下载")
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZEnd)
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x}) //原来是console打印
								// mySSh.Websocket.WriteJSON(&WebMsg{Type: "console", ReslutData: string(sshOut[:n])})
							}
						} else if _, ok := IsContain(sshOut[:n], ZModemCancel); ok { //判断是否取消下载
							fmt.Println("取消下载")
							s.ZModemSZ = false
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, sshOut[:n])
						} else { //都没有直接发送下载的数据
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, sshOut[:n])
						}

					} else if s.ZModemRZ {
						// fmt.Println("rz")
						if x, ok := IsContain(sshOut[:n], ZModemRZEnd); ok {
							s.ZModemRZ = false
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemRZEnd)
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})

							}
						} else if _, ok := IsContain(sshOut[:n], ZModemCancel); ok {
							s.ZModemRZ = false
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, sshOut[:n])
							fmt.Println("ssh-->web" + string(sshOut[:n]))
						} else {
							// rz 上传过程中服务器端还是会给客户端发送一些信息，比如心跳
							//ws.Ws.WriteJSON(&message{Type: messageTypeConsole, Data: sshOut[:n]})
							//ws.Ws.WriteMessage(websocket.BinaryMessage, sshOut[:n])
							startIndex := bytes.Index(sshOut[:n], ZModemRZCtrlStart)
							if startIndex != -1 {
								endIndex := bytes.Index(sshOut[:n], ZModemRZCtrlEnd1)
								if endIndex != -1 {
									ctrl := append(ZModemRZCtrlStart, sshOut[startIndex+len(ZModemRZCtrlStart):endIndex]...)
									ctrl = append(ctrl, ZModemRZCtrlEnd1...)
									_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ctrl)
									info := append(sshOut[:startIndex], sshOut[endIndex+len(ZModemRZCtrlEnd1):n]...)
									if len(info) != 0 {
										_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: info})
									}
								} else {
									endIndex = bytes.Index(sshOut[:n], ZModemRZCtrlEnd2)
									if endIndex != -1 {
										ctrl := append(ZModemRZCtrlStart, sshOut[startIndex+len(ZModemRZCtrlStart):endIndex]...)
										ctrl = append(ctrl, ZModemRZCtrlEnd2...)
										_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ctrl)
										info := append(sshOut[:startIndex], sshOut[endIndex+len(ZModemRZCtrlEnd2):n]...)
										if len(info) != 0 {
											_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: info})
										}
									} else {
										_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: sshOut[:n]})
									}
								}
							} else {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: sshOut[:n]})
							}
						}
					} else {
						if x, ok := IsContain(sshOut[:n], ZModemSZStart); ok { //是否包含sz

							if y, ok := IsContain(x, ZModemCancel); ok {
								// 下载不存在的文件以及文件夹(zmodem 不支持下载文件夹)时
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: y})
								// mySSh.Websocket.WriteJSON(&WebMsg{Type: "stdin", ReslutData: string(y)})

							} else {
								s.ZModemSZ = true
								if len(x) != 0 {
									_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})
									// mySSh.Websocket.WriteJSON(&WebMsg{Type: "console", ReslutData: string(x)})
								}
								_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemSZStart)
							}

						} else if x, ok := IsContain(sshOut[:n], ZModemRZStart); ok {

							s.ZModemRZ = true
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})
							}
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemRZStart)

						} else if x, ok := IsContain(sshOut[:n], ZModemRZEStart); ok {

							s.ZModemRZ = true
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})
								// mySSh.Websocket.WriteJSON(&WebMsg{Type: "console", ReslutData: string(x)})
							}
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemRZEStart)

						} else if x, ok := IsContain(sshOut[:n], ZModemRZSStart); ok {

							s.ZModemRZ = true
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})
								// mySSh.Websocket.WriteJSON(&WebMsg{Type: "console", ReslutData: string(x)})
							}
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemRZSStart)

						} else if x, ok := IsContain(sshOut[:n], ZModemRZESStart); ok {

							s.ZModemRZ = true
							if len(x) != 0 {
								_ = s.Websocket.WriteJSON(&WebMsg{Type: "console", Data: x})
								// mySSh.Websocket.WriteJSON(&WebMsg{Type: "console", ReslutData: string(x)})
							}
							_ = s.Websocket.WriteMessage(websocket.BinaryMessage, ZModemRZESStart)

						} else {
							// 如果不满足上面所有的条件 说明只是普通的命令回执显示，直接发回前	端即可
							//// 判断字符类型
							//if IsGBK(gbkOut) {
							//	gbkOut, err = simplifiedchinese.GB18030.NewDecoder().Bytes(gbkOut)
							//	if err != nil {
							//		gbkOut = sshOut[:n]
							//	}
							//}
							_ = s.Websocket.WriteJSON(&WebMsg{Type: "stdin", Data: gbkOut})
							// 将操作记录到文本中
							// 写入操作的时间线
							_, _ = fileTimeLine.Write([]byte(fmt.Sprintf("%d\t%d\n", time.Now().UnixMilli(), len(gbkOut))))
							//写入操作的内容
							_, _ = fileOperator.Write(gbkOut)
							// 高危命令判定前，把用户输入的内容存起来留作判定
						}
					}
				}
			}

		}

	}
}

func (s *Shell) Stop() {
	defer func() {
		recover()
	}()

	<-s.Ctx.Done()
	go func() {
		time.Sleep(5 * time.Second)
		// 计算文件MD5值
		fileInfo, err := dir.GetFileInfo(path.Join(s.FilePath, s.LogShell.Filename+".data"))
		// 计算文件大小
		//	 保存到数据库中
		_, err = dao.LogShell.Ctx(context.Background()).Where(dao.LogShell.Columns().Id, s.LogShell.Id).Update(do.LogShell{
			Online:   false,
			Size:     fileInfo.Size,
			Md5:      fmt.Sprintf("%x", fileInfo.MD5),
			LogoutAt: time.Now().Unix(),
		})
		if err != nil {
			g.Log().Error(s.Ctx, err.Error())
		}
	}()
	// 最后释放，避免panic
	_ = s.Websocket.Close()
	_ = s.Pty.Close()

}
