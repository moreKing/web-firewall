package public

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/service"
	"time"

	"server/api/public/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Session().RemoveOnlineUser(ctx, ghttp.RequestFromCtx(ctx).GetHeader("Authentication"), "手动退出", time.Now().Unix())
	return
}
