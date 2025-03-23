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

func (c *ControllerV1) DeleteDnatPolicy(ctx context.Context, req *v1.DeleteDnatPolicyReq) (res *v1.DeleteDnatPolicyRes, err error) {
	err = dao.DnatRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.DnatRules{}).Where(dao.DnatRules.Columns().Id, req.ID).Delete()
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
