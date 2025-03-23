package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetMessage(ctx context.Context, req *v1.SetMessageReq) (res *v1.SetMessageRes, err error) {
	g.Log().Debug(ctx, "call v1.SetMessage", req)
	return nil, service.SystemConfig().SetMessage(req.Message)
}
