package audit

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/service"

	"server/api/audit/v1"
)

func (c *ControllerV1) GetShellReplayToken(ctx context.Context, req *v1.GetShellReplayTokenReq) (res *v1.GetShellReplayTokenRes, err error) {
	m := service.BizCtx().Get(ctx)
	// 获取一个临时token 过期时间1分钟
	token, err := service.CodeServer().CreateWebsocketToken(ctx, m.User.ID, ghttp.RequestFromCtx(ctx).GetClientIp(), 1)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &v1.GetShellReplayTokenRes{Token: token.Token}, nil
}
