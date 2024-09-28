package policy

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"

	"server/api/policy/v1"
)

func (c *ControllerV1) ReplaceOutputPolicy(ctx context.Context, req *v1.ReplaceOutputPolicyReq) (res *v1.ReplaceOutputPolicyRes, err error) {

	var rule entity.OutputRules
	err = dao.OutputRules.Ctx(ctx).Where(dao.OutputRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.OutputRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.OutputRules{}).Where(dao.OutputRules.Columns().Id, req.ID).Update(do.OutputRules{
			Protocol: req.Protocol,
			Port:     req.Port,
			Ip:       req.Ip,
			Ct:       req.Ct,
			Icmp:     req.Icmp,
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
