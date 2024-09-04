package audit

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"
	"time"

	"server/api/audit/v1"
)

func (c *ControllerV1) CutLogin(ctx context.Context, req *v1.CutLoginReq) (res *v1.CutLoginRes, err error) {
	m := service.BizCtx().Get(ctx)
	//user, err := service.User().GetUserByID(ctx, m.User.ID)
	user := model.WithUser{}
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, m.User.ID).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	//
	desUser, ok := service.Session().GetOnlineUser(req.UUID)
	g.Log().Debug(ctx, "onlineUsers ok: ", ok)
	if !ok {
		return nil, errors.New("不存在的在线token")
	}

	dbuser := model.WithUser{}
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, desUser.ID).Scan(&dbuser)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// 删除
	service.Session().RemoveOnlineUser(ctx, req.UUID, fmt.Sprintf("%s 切断登录", user.Loginname), time.Now().Unix())

	return nil, nil
}
