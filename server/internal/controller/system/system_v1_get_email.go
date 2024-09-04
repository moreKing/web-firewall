package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) GetEmail(ctx context.Context, req *v1.GetEmailReq) (res *v1.GetEmailRes, err error) {
	return &v1.GetEmailRes{Email: service.SystemConfig().GetEmail()}, nil
}
