package policy

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"server/api/policy/v1"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
)

func (c *ControllerV1) ReplaceOutputLimitPolicy(ctx context.Context, req *v1.ReplaceOutputLimitPolicyReq) (res *v1.ReplaceOutputLimitPolicyRes, err error) {

	var rule entity.OutputLimitRules
	err = dao.OutputLimitRules.Ctx(ctx).Where(dao.OutputLimitRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.OutputLimitRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.OutputLimitRules{}).Where(dao.OutputLimitRules.Columns().Id, req.ID).Update(do.OutputLimitRules{
			Protocol: req.Protocol,
			Port:     req.Port,
			Ip:       req.Ip,
			Comment:  req.Comment,
			Limit:    req.Limit,
			Speed:    req.Speed,
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
