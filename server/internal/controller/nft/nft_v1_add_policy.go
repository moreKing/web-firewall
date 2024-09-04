package nft

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/model"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"server/api/nft/v1"
)

func (c *ControllerV1) AddPolicy(ctx context.Context, req *v1.AddPolicyReq) (res *v1.AddPolicyRes, err error) {
	err = service.Nft().Add(ctx, &model.Rulesets{
		Meta: gmeta.Meta{},
		Rulesets: &entity.Rulesets{
			Comment:   req.Comment,
			Position:  req.Position,
			Chain:     req.Chain,
			CreatedAt: time.Now().Unix(),
		},
		Handle: "",
		Expr:   req.Expr,
	}, req.Add)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
