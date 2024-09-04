package nft

import (
	"context"
	"server/internal/service"
	"time"

	"server/api/nft/v1"
)

func (c *ControllerV1) GetPolicy(ctx context.Context, req *v1.GetPolicyReq) (res *v1.GetPolicyRes, err error) {
	return &v1.GetPolicyRes{
		Data:      *(service.Nft().GetChainList(req.Chain)),
		Total:     len(*(service.Nft().GetChainList(req.Chain))),
		Timestamp: time.Now().Unix(),
	}, nil
}
