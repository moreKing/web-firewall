package nft

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/service"

	"server/api/nft/v1"
)

func (c *ControllerV1) DeletePolicy(ctx context.Context, req *v1.DeletePolicyReq) (res *v1.DeletePolicyRes, err error) {

	err = service.Nft().Delete(ctx, req.ID)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return nil, nil
}
