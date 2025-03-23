package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetWebConfig(ctx context.Context, req *v1.SetWebConfigReq) (res *v1.SetWebConfigRes, err error) {
	return nil, service.SystemConfig().SetWebSession(req.WebSession)
}
