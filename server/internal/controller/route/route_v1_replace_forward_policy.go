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

func (c *ControllerV1) ReplaceForwardPolicy(ctx context.Context, req *v1.ReplaceForwardPolicyReq) (res *v1.ReplaceForwardPolicyRes, err error) {
	var rule entity.ForwardRules
	err = dao.ForwardRules.Ctx(ctx).Where(dao.ForwardRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.ForwardRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.ForwardRules{}).Where(dao.ForwardRules.Columns().Id, req.ID).Update(do.ForwardRules{
			Protocol: req.Protocol,
			Sip:      req.Sip,
			Dip:      req.Dip,
			Port:     req.Port,
			Comment:  req.Comment,
			Policy:   req.Policy,
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
