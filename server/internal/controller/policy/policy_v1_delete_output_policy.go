package policy

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/service"

	"server/api/policy/v1"
)

func (c *ControllerV1) DeleteOutputPolicy(ctx context.Context, req *v1.DeleteOutputPolicyReq) (res *v1.DeleteOutputPolicyRes, err error) {
	err = dao.OutputRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.OutputRules{}).Where(dao.OutputRules.Columns().Id, req.ID).Delete()
		if err != nil {
			return err
		}

		return service.Policy().Flush(ctx)
	})

	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return nil, nil
}
