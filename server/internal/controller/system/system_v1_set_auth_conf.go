package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetAuthConf(ctx context.Context, req *v1.SetAuthConfReq) (res *v1.SetAuthConfRes, err error) {
	err = service.SystemConfig().SetAuthConf(ctx, req.AuthenticateConf)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return nil, nil
}
