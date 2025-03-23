package route

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"

	"server/api/route/v1"
)

func (c *ControllerV1) ReplaceSnatPolicy(ctx context.Context, req *v1.ReplaceSnatPolicyReq) (res *v1.ReplaceSnatPolicyRes, err error) {

	var rule entity.SnatRules
	err = dao.SnatRules.Ctx(ctx).Where(dao.SnatRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.SnatRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.SnatRules{}).Where(dao.SnatRules.Columns().Id, req.ID).Update(do.SnatRules{
			Sip:     req.Sip,
			Dip:     req.Dip,
			Oif:     req.Oif,
			Snat:    req.Snat,
			Comment: req.Comment,
		})

		if err != nil {
			return err
		}
		return service.Policy().Flush(ctx)
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
