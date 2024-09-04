package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"regexp"
	"server/internal/model"
	"server/internal/service"
	"time"
)

var regMask = regexp.MustCompile(`"[a-zA-Z]*[Pp]assword":"([^"]*)"`)
var regMaskContent = regexp.MustCompile(`:"([^"]*)"`)

type sMiddleware struct {
	onlineUsers map[string]*model.ContextUser
}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{
		onlineUsers: make(map[string]*model.ContextUser),
	}
}

// Ctx 将自定义业务上下文变量注入到当前请求的上下文中
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			ID:        user.Id,
			LoginName: user.Loginname,
			Password:  user.Password,
		}
	}
	//Continue execution of next middleware.
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing. 提供前端跨域（不开启，前端自行代理解决）
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// Auth validates the request to allow only signed-in users visit. 校验登录状态是否可用
func (s *sMiddleware) Auth(r *ghttp.Request) {
	//g.Log().Debug(r.Context(), "cookie", r.Cookie.GetSessionId())
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
