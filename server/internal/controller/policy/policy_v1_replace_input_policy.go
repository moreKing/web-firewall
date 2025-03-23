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

func (c *ControllerV1) ReplaceInputPolicy(ctx context.Context, req *v1.ReplaceInputPolicyReq) (res *v1.ReplaceInputPolicyRes, err error) {

	var rule entity.InputRules
	err = dao.InputRules.Ctx(ctx).Where(dao.InputRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.InputRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.InputRules{}).Where(dao.InputRules.Columns().Id, req.ID).Update(do.InputRules{
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
