package route

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/service"

	"server/api/route/v1"
)

func (c *ControllerV1) DeleteForwardPolicy(ctx context.Context, req *v1.DeleteForwardPolicyReq) (res *v1.DeleteForwardPolicyRes, err error) {
	err = dao.ForwardRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.ForwardRules{}).Where(dao.ForwardRules.Columns().Id, req.ID).Delete()
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
