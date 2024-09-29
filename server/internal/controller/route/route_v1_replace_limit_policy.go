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

func (c *ControllerV1) ReplaceLimitPolicy(ctx context.Context, req *v1.ReplaceLimitPolicyReq) (res *v1.ReplaceLimitPolicyRes, err error) {
	var rule entity.ForwardLimitRules
	err = dao.ForwardLimitRules.Ctx(ctx).Where(dao.ForwardLimitRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.ForwardLimitRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.ForwardLimitRules{}).Where(dao.ForwardLimitRules.Columns().Id, req.ID).Update(do.ForwardLimitRules{
			Protocol: req.Protocol,
			PortType: req.PortType,
			Port:     req.Port,
			Sip:      req.Sip,
			Dip:      req.Dip,
			Limit:    req.Limit,
			Speed:    req.Speed,

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
