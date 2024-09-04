package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) GetMessage(ctx context.Context, req *v1.GetMessageReq) (res *v1.GetMessageRes, err error) {
	return &v1.GetMessageRes{
		Message: service.SystemConfig().GetMessage(),
	}, nil
}
