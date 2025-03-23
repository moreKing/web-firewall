package policy

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/service"

	"server/api/policy/v1"
)

func (c *ControllerV1) DeleteInputPolicy(ctx context.Context, req *v1.DeleteInputPolicyReq) (res *v1.DeleteInputPolicyRes, err error) {

	if req.ID == 4 {
		return nil, errors.New("不能删除内置web端口策略")
	}

	err = dao.InputRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.InputRules{}).Where(dao.InputRules.Columns().Id, req.ID).Delete()
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
