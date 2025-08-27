package system

import (
	"context"
	v1 "server/api/system/v1"
	"server/internal/model"
	"server/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SetEmail(ctx context.Context, req *v1.SetEmailReq) (res *v1.SetEmailRes, err error) {
	g.Log().Debug(ctx, "SetEmail ----------------------------- ", req)
	if err := service.SystemConfig().SetEmail(&model.Email{
		Enable:   req.Enable,
		SMTP:     req.SMTP,
		Port:     req.Port,
		Email:    req.Email,
		Account:  req.Account,
		Protocol: req.Protocol,
		Password: req.Password,
	}); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return nil, nil
}
