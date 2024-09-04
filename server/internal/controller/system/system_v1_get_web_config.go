package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) GetWebConfig(ctx context.Context, req *v1.GetWebConfigReq) (res *v1.GetWebConfigRes, err error) {

	return &v1.GetWebConfigRes{WebSession: service.SystemConfig().GetWebSession()}, nil
}
