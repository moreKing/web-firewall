package public

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/dao"
	"server/internal/service"
	"server/utility/gm"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/public/v1"
)

func (c *ControllerV1) SetUserPassword(ctx context.Context, req *v1.SetUserPasswordReq) (res *v1.SetUserPasswordRes, err error) {
	// 校验原密码是否正确
	user, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, service.BizCtx().Get(ctx).User.ID).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	if gm.SM3(req.OldPassword, user[dao.Users.Columns().Slat].String()) != user[dao.Users.Columns().Password].String() {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "旧密码校验失败")
	}

	// 调用用户自己的属性去更新密码
	err = service.User().SetNativePassword(ctx, service.BizCtx().Get(ctx).User.ID, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &v1.SetUserPasswordRes{}, nil
}
