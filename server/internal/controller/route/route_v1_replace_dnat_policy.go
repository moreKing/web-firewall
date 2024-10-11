package route

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"

	"server/api/route/v1"
)

func (c *ControllerV1) ReplaceDnatPolicy(ctx context.Context, req *v1.ReplaceDnatPolicyReq) (res *v1.ReplaceDnatPolicyRes, err error) {

	if len(req.Port) < 1 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "端口映射不能为空")
	}

	var rule entity.DnatRules
	err = dao.DnatRules.Ctx(ctx).Where(dao.DnatRules.Columns().Id, req.ID).Scan(&rule)
	if err != nil {
		return nil, err
	}

	err = dao.DnatRules.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		_, err := tx.Ctx(ctx).Model(&do.DnatRules{}).Where(dao.DnatRules.Columns().Id, req.ID).Update(do.DnatRules{

			Protocol: "",
			Dip:      req.Dip,
			Iif:      req.Iif,
			Port:     req.Port,
			Dnat:     req.Dnat,
			Comment:  req.Comment,
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
