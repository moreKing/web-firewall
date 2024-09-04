package nft

import (
	"context"
	"server/internal/service"

	"server/api/nft/v1"
)

func (c *ControllerV1) ChangePolicyPosition(ctx context.Context, req *v1.ChangePolicyPositionReq) (res *v1.ChangePolicyPositionRes, err error) {
	// 修改策略位置

	if req.ID == req.Position {
		return
	}

	err = service.Nft().UpdatePosition(ctx, req.ID, req.Position, req.Add)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
