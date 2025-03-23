package login

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xlzd/gotp"
	"server/api/login/v1"
	"server/internal/dao"
	"server/internal/global"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility/gm"
	"strings"
	"time"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	isSave := true
	// 创建一条登录记录
	logs := entity.LogLogins{
		Uuid:      gonanoid.Must(),
		Loginname: req.Username,
		ClientIp:  ghttp.RequestFromCtx(ctx).GetClientIp(),
		UserId:    0,
		TotpCode:  "",
		Success:   false,
		Online:    false,
		Log:       "",
		LoginAt:   time.Now().Unix(),
		LogoutAt:  0,
	}

	// 保存到数据库
	defer func() {
		if isSave {
			logs.LogoutAt = time.Now().Add(time.Duration(service.SystemConfig().GetWebSession().Timeout) * time.Minute).Unix()
			_, err2 := dao.LogLogins.DB().Insert(ctx, dao.LogLogins.Table(), logs)
			if err2 != nil {
				g.Log().Error(ctx, err2)
			}

			//	如果登录成功 更新用户信息
			if logs.Success {
				// 添加在线用户
				service.Session().AddOnlineUser(&model.ContextUser{
					ID:             logs.UserId,
					Authentication: logs.Uuid,
					LoginName:      logs.Loginname,
					Password:       req.Password,
					ClientIP:       logs.ClientIp,
					CreatedAt:      logs.LoginAt,
					ExpiredAt:      logs.LogoutAt,
				})
				// 写入返回cookie
				ghttp.RequestFromCtx(ctx).Cookie.SetSessionId(logs.Uuid)
				res = &v1.LoginRes{Authentication: logs.Uuid}
				// 更新用户
				_, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, logs.UserId).Fields(dao.Users.Columns().LastloginAt).Update(do.Users{LastloginAt: time.Now().Unix()})
				if err != nil {
					g.Log().Error(ctx, err2)
				}
			}
		}
	}()

	// 第一步 需要到数据库查询有没有这个用户
	var user *model.WithUser
	err = dao.Users.Ctx(ctx).WithAll().Where(dao.Users.Columns().Loginname, req.Username).Scan(&user)
	if err != nil || user == nil {
		logs.Log = "账号不存在"
		g.Log().Error(ctx, logs.Log, err)
		return nil, gerror.New("账号或密码错误")
	}
	logs.UserId = user.Id
	logs.Username = user.Username

	//fmt.Printf("%#v\n", user.Authenticates)
	//g.Log().Debug(ctx, user)

	//1. 判断密码是否正确
	if gm.SM3(req.Password, user.Slat) != user.Password {
		logs.Log = "密码错误"
		g.Log().Error(ctx, logs.Log)
		return nil, gerror.New("账号或密码错误")
	}

	g.Log().Debug(ctx, user)

	// 不是双因素认证 1 手机令牌 2 邮件 3 短信
	if user.AuthenticateId == 0 {
		// 登录成功
		logs.Success = true
		logs.Online = true
		return
	}

	// 是双因素认证的情况，要看到底是哪种认证方式
	// code 特殊返回 1 totp未绑定 2 需要输入2步令牌
	switch user.AuthenticateId {
	case 1:

		if !user.TotpState && req.Code != "" {
			logs.TotpCode = req.Code
			// 校验是否正确
			if err := service.CodeServer().VerifyTotp(ctx, req.Code, user.TotpToken, user.Id); err != nil {
				logs.Log = err.Error()
				g.Log().Error(ctx, logs.Log)
				return nil, gerror.New("账号或密码错误")
			}

			// 登录成功
			logs.Success = true
			logs.Online = true
			return

		}

		//	手机令牌 当前用户是否需要绑定手机令牌
		if !user.TotpState && req.Code == "" {
			isSave = false
			return nil, gerror.NewCode(gcode.New(2, "请输入二步令牌", nil))
		}

		//	首次绑定令牌
		if user.TotpState && req.Code != "" {
			logs.TotpCode = req.Code
			// 校验令牌是否正确
			if err := service.CodeServer().VerifyTotp(ctx, req.Code, user.TotpToken, user.Id); err != nil {
				logs.Log = err.Error()
				g.Log().Error(ctx, logs.Log)
				return nil, gerror.New("账号或密码错误")
			}

			//将用户 totpState 改成false
			_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, logs.UserId).Fields(dao.Users.Columns().TotpState).Update(do.Users{TotpState: false})
			if err != nil {
				logs.Log = err.Error()

				g.Log().Error(ctx, "修改用户totp state 发生错误 ", err)
				return nil, gerror.New("账号或密码错误")
			}

			// 登录成功
			logs.Success = true
			logs.Online = true
			return
		}

		if user.TotpState && req.Code == "" {
			//logs.Log = "返回信息让用户绑定手机令牌"
			otpToken := gotp.RandomSecret(20)
			if otpToken == "" {
				logs.Log = "无法生成手机令牌token进行绑定"
				g.Log().Error(ctx, logs.Log)
				return nil, gerror.New("发生未知错误，请管理员查看登录日志")
			}

			// 生成成功，保存到数据库中
			_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, user.Id).Fields(dao.Users.Columns().TotpToken).Update(do.Users{TotpToken: otpToken})
			if err != nil {
				logs.Log = err.Error()
				g.Log().Error(ctx, err)
				return nil, gerror.New("发生未知错误，请管理员查看登录日志")
			}

			isSave = false
			g.Log().Debug(ctx, "返回信息让用户绑定手机令牌", service.SystemConfig().GetAuthConf().TotpIssuer)
			return &v1.LoginRes{
				Authentication: otpToken,
				Issuer:         service.SystemConfig().GetAuthConf().TotpIssuer,
			}, gerror.NewCode(gcode.New(1, "", nil))

		}

	case 2:
		//	发邮件
		if req.Code == "" {
			isSave = false
			// 创建验证码
			code, err := service.CodeServer().CreateCode(ctx, user.Id, ghttp.RequestFromCtx(ctx).GetClientIp(), service.SystemConfig().GetAuthConf().EmailOffset)
			if err != nil {
				g.Log().Error(ctx, err)
				logs.Log = err.Error()
				return nil, gerror.New("服务错误，请联系管理员查看")
			}

			content, err := g.View().ParseContent(
				ctx, global.SendEmailCodeTpl,
				g.Map{"username": user.Username,
					"account": service.SystemConfig().GetEmail().Account,
					"code":    code.Code,
					"offset":  service.SystemConfig().GetAuthConf().MessageOffset},
			)

			if err != nil {
				g.Log().Error(ctx, err)
				logs.Log = err.Error()
				return nil, gerror.New("服务错误，请联系管理员查看")
			}

			// 发送短信
			err = service.SystemConfig().SendEmailHtml(user.Email, "登录验证码", content)
			if err != nil {
				g.Log().Error(ctx, err)
				logs.Log = err.Error()
				return nil, gerror.New("服务错误，请联系管理员查看")
			}

			return &v1.LoginRes{Authentication: code.Token}, gerror.NewCode(gcode.New(2, "请输入二步令牌", nil))
		}

		if req.Code != "" && strings.TrimSpace(req.Token) != "" {
			if !service.CodeServer().VerifyCode(ctx, user.Id, req.Token, req.Code, ghttp.RequestFromCtx(ctx).GetClientIp()) {
				// 没有通过校验
				logs.Log = "邮件验证失败"
				g.Log().Error(ctx, logs.Log)
				return nil, gerror.New("账号或密码错误")
			}
			// 登录成功
			logs.Success = true
			logs.Online = true
			return
		}

		if req.Code != "" && strings.TrimSpace(req.Token) == "" {
			logs.Log = "疑似攻击：携带code 没有携带token"
			g.Log().Error(ctx, logs.Log)
			return nil, gerror.New("账号或密码错误")
		}

	case 3:
		//	发短信
		if req.Code == "" {
			isSave = false
			// 创建验证码
			code, err := service.CodeServer().CreateCode(ctx, user.Id, ghttp.RequestFromCtx(ctx).GetClientIp(), service.SystemConfig().GetAuthConf().EmailOffset)
			if err != nil {
				g.Log().Error(ctx, err)
				logs.Log = err.Error()
				return nil, gerror.New("服务错误，请联系管理员查看")
			}

			// 发送短信
			err = service.SystemConfig().SendMessage(ctx, code.Code, user.Mobile)
			if err != nil {
				g.Log().Error(ctx, err)
				logs.Log = err.Error()
				return nil, gerror.New("服务错误，请联系管理员查看")
			}

			return &v1.LoginRes{Authentication: code.Token}, gerror.NewCode(gcode.New(2, "请输入二步令牌", nil))
		}

		if req.Code != "" && strings.TrimSpace(req.Token) != "" {
			if !service.CodeServer().VerifyCode(ctx, user.Id, req.Token, req.Code, ghttp.RequestFromCtx(ctx).GetClientIp()) {
				// 没有通过校验
				logs.Log = "短信验证失败"
				g.Log().Error(ctx, logs.Log)
				return nil, gerror.New("账号或密码错误")
			}
			// 登录成功
			logs.Success = true
			logs.Online = true
			return
		}

		if req.Code != "" && strings.TrimSpace(req.Token) == "" {
			logs.Log = "疑似攻击：携带code 没有携带token"
			g.Log().Error(ctx, logs.Log)
			return nil, gerror.New("账号或密码错误")
		}
	default:
		logs.Log = fmt.Sprintf("未知的二步认证方式：%d", user.AuthenticateId)
		g.Log().Error(ctx, logs.Log)
		return nil, gerror.New("账号或密码错误")

	}

	return
}
