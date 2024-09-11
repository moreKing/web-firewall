package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"server/internal/model"
	"server/internal/service"
	"strings"
	"time"
)

type req struct {
	Method string
	Url    string
}

var whiteList = []req{
	{
		Method: "get",
		Url:    "/api/v1/audit/shell-replay",
	},
}

func isWhiteList(method, url string) bool {
	for _, r := range whiteList {
		if strings.ToUpper(r.Method) == strings.ToUpper(method) && strings.ToUpper(r.Url) == strings.ToUpper(url) {
			return true
		}
	}
	return false
}

// Auth  校验登录状态是否可用
func (s *sMiddleware) Auth(r *ghttp.Request) {

	if isWhiteList(r.Method, r.Request.URL.Path) {
		// 校验一下token
		userid, err := service.CodeServer().VerifyWebsocketToken(r.Context(), r.Request.URL.Query().Get("token"), r.GetClientIp())
		if err != nil {
			r.Response.WriteStatus(http.StatusUnauthorized, "无效token")
			return
		}

		customCtx := &model.Context{
			Session: r.Session,
			User: &model.ContextUser{
				ID: userid,
			},
		}

		service.BizCtx().Init(r, customCtx)
		r.Middleware.Next()
		return
	}

	// cookie 与 Authentication 满足一个即可 Authentication优先级高于cookie
	Authentication := r.GetHeader("Authentication")
	if r.Method == "GET" && r.Request.URL.Path == "/api/v1/system/shell" {
		Authentication = r.Request.URL.Query().Get("Authorization")
	}
	g.Log().Debug(r.Context(), "Header Authentication ", Authentication)
	if Authentication == "" {
		Authentication = r.Cookie.GetSessionId()
		//	注入到Header里后续中间件统一
		r.Header.Set("Authentication", Authentication)
	}
	g.Log().Debug(r.Context(), "Final Authentication ", Authentication)
	if Authentication != "" {
		user, ok := service.Session().GetOnlineUser(Authentication)
		g.Log().Debug(r.Context(), "onlineUsers ok: ", ok)
		if !ok {
			r.Response.WriteStatus(http.StatusUnauthorized, "登录验证失败")
			return
		}

		if user != nil && user.ClientIP == r.GetClientIp() && time.Now().Unix() < user.ExpiredAt {
			g.Log().Debug(r.Context(), "authorized ok")
			user.ExpiredAt = time.Now().Add(time.Duration(service.SystemConfig().GetWebSession().Timeout) * time.Minute).Unix()
			// 校验通过 往里面注入userid 方便后续获取用户信息呢
			//service.Middleware().Ctx(r)
			customCtx := &model.Context{
				Session: r.Session,
				User: &model.ContextUser{
					ID:        user.ID,
					LoginName: user.LoginName,
					Password:  user.Password,
				},
			}
			service.BizCtx().Init(r, customCtx)

			r.Middleware.Next()
			return
		}
		g.Log().Debug(r.Context(), "onlineUsers: ", user)

	}
	r.Response.WriteStatus(http.StatusUnauthorized, "登录验证失败")
}
