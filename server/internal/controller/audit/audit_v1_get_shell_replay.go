package audit

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"path"
	"server/internal/dao"
	"server/internal/model/entity"
	"strconv"
	"strings"
	"time"

	"server/api/audit/v1"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebMsg struct {
	Type string `json:"type"`
	Cols int    `json:"Cols"`
	Rows int    `json:"rows"`
	Data []byte `json:"data"`
}

func (c *ControllerV1) GetShellReplay(ctx context.Context, req *v1.GetShellReplayReq) (res *v1.GetShellReplayRes, err error) {

	// 回放相关
	var log entity.LogShell
	err = dao.LogShell.Ctx(ctx).Where(dao.LogShell.Columns().Id, req.Id).Scan(&log)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	webcon, err := upGrader.Upgrade(ghttp.RequestFromCtx(ctx).Response.Writer, ghttp.RequestFromCtx(ctx).Request, nil)
	if err != nil {
		//fmt.Println("升级http 为websoket失败：", err)
		g.Log().Error(ctx, "升级http 为websoket失败：", err)
		return nil, errors.New("websocket建立失败")
	}
	defer func(webcon *websocket.Conn) {
		_ = webcon.Close()

	}(webcon)

	g.Log().Debug(ctx, "websocket建立成功")
	// 读取审计文件
	logPath := g.Cfg().MustGet(ctx, "server.shellLog", "/var/log/web-firewall").String()

	createTime := time.UnixMilli(log.CreatedAt * 1000)

	fullPath := path.Join(logPath, fmt.Sprintf("%d/%d/%d", createTime.Year(), createTime.Month(), createTime.Day()), log.Filename)
	tl, err := os.ReadFile(fullPath + ".tl")
	if err != nil {
		fmt.Println("读取时间线文件失败：", err)
		return
	}
	// string(tl)
	lines := strings.Split(string(tl), "\n")
	//解决index out问题
	lines = lines[:(len(lines) - 1)]

	//打开操作记录文件
	auditFile, err := os.Open(fullPath + ".data")
	if err != nil {
		fmt.Println("读取操作记录文件文件失败：", err)
		return
	}
	g.Log().Debug(ctx, "lines个数： %d", len(lines))

	var lastTime = 0

	for _, line := range lines {
		lens := strings.Split(line, "\t")
		g.Log().Debug(ctx, lens)
		tmpTime, err1 := strconv.Atoi(strings.TrimSpace(lens[0]))
		if err1 != nil {
			g.Log().Error(ctx, "时间戳转换成数字失败", err1)
			return nil, err1
		}

		bufferSize, err1 := strconv.Atoi(strings.TrimSpace(lens[1]))
		if err1 != nil {
			g.Log().Error(ctx, "长度转换成数字失败", err1)
			return nil, err1
		}

		buffer := make([]byte, bufferSize)
		_, err = auditFile.Read(buffer)
		if err != nil {
			g.Log().Error(ctx, "for 读取操作记录文件文件失败：", err)
			return nil, err
		}

		_ = webcon.WriteJSON(&WebMsg{Type: "stdin", Data: buffer})

		if tmpTime-lastTime > 500 {
			time.Sleep(500 * time.Millisecond)
			g.Log().Debug(ctx, "等待时间： 1s")
		} else {
			time.Sleep(time.Duration(tmpTime-lastTime) * time.Millisecond)
			g.Log().Debug(ctx, "等待时间： ", tmpTime-lastTime)
		}

		lastTime = tmpTime
	}

	return nil, nil
}
