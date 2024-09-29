package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"server/internal/service"
)

func (s *sMiddleware) IsForward(r *ghttp.Request) {

	kernel, err := service.Kernel().Get(r.Context())
	if err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.WriteStatus(http.StatusInternalServerError, err)
		return
	}

	if !kernel.Forward {
		r.Response.WriteJsonExit(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    581,
			Message: "Forward未启用",
		})
		return
	}
	r.Middleware.Next()
}
