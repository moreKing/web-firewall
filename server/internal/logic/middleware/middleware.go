package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"regexp"
	"server/internal/model"
	"server/internal/service"
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
