package home

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/global"
	"server/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/home/v1"
)

func (c *ControllerV1) GetHome(ctx context.Context, req *v1.GetHomeReq) (res *v1.GetHomeRes, err error) {

	online, err := dao.LogLogins.Ctx(ctx).Where(dao.LogLogins.Columns().Online, true).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "获取在线用户数量失败")
	}

	//

	return &v1.GetHomeRes{
		Home: &model.Home{
			Assets:    0,
			Users:     0,
			Online:    online,
			Accounts:  0,
			Yesterday: *global.GetYesterday(ctx),
			Today:     *global.GetToday(ctx),
			License:   0,
		},
	}, nil
}
