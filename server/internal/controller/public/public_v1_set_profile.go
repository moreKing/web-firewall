package public

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/service"

	"server/api/public/v1"
)

func (c *ControllerV1) SetProfile(ctx context.Context, req *v1.SetProfileReq) (res *v1.SetProfileRes, err error) {

	// 调用用户自己的属性去更新密码
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, service.BizCtx().Get(ctx).User.ID).Update(&do.Users{
		Meta:           g.Meta{},
		Username:       req.Username,
		TotpState:      req.Bind,
		Email:          req.Email,
		Mobile:         req.Phone,
		AuthenticateId: req.AuthenticateId,
	})

	if err != nil {
		return nil, err
	}
	return &v1.SetProfileRes{}, nil
}
