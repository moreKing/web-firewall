package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) TestMessage(ctx context.Context, req *v1.TestMessageReq) (res *v1.TestMessageRes, err error) {
	err = service.SystemConfig().SendMessage(ctx, "测试内容", req.To)
	return nil, err
}
