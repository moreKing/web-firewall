package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) GetAuthConf(ctx context.Context, req *v1.GetAuthConfReq) (res *v1.GetAuthConfRes, err error) {
	return &v1.GetAuthConfRes{AuthenticateConf: service.SystemConfig().GetAuthConf()}, nil
}
