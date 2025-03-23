package public

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/global"
	"server/internal/model"
	"server/internal/service"

	"server/api/public/v1"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	m := service.BizCtx().Get(ctx)
	g.Log().Debug(ctx, "user初始化 ", m)
	user := &model.UserInfo{}
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, m.User.ID).WithAll().FieldsEx(dao.Users.Columns().Slat, dao.Users.Columns().Password, dao.Users.Columns().TotpToken).Scan(user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	g.Log().Debug(ctx, user)
	res = &v1.ProfileRes{
		UserInfo:    user,
		Routes:      global.GetAllRoutes(), // 为后面多用户准备
		Home:        "home",                // 为后面多用户准备
		Permissions: &[]string{},           // 为后面多用户准备
		Email:       service.SystemConfig().GetEmail().Enable,
		Message:     service.SystemConfig().GetMessage().State > 0,
	}

	return
}
