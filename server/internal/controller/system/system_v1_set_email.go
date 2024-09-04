package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetEmail(ctx context.Context, req *v1.SetEmailReq) (res *v1.SetEmailRes, err error) {
	return nil, service.SystemConfig().SetEmail(req.Email)
}
