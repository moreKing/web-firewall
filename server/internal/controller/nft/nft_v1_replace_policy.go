package nft

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"

	"server/api/nft/v1"
)

func (c *ControllerV1) ReplacePolicy(ctx context.Context, req *v1.ReplacePolicyReq) (res *v1.ReplacePolicyRes, err error) {

	var rule model.Rulesets
	err = dao.Rulesets.Ctx(ctx).Where(dao.Rulesets.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	rule.Comment = req.Comment

	err = service.Nft().Replace(ctx, &model.Rulesets{
		Rulesets: rule.Rulesets,
		Expr:     req.Expr,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return nil, nil
}
