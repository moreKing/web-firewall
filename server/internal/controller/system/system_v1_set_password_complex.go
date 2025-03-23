package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetPasswordComplex(ctx context.Context, req *v1.SetPasswordComplexReq) (res *v1.SetPasswordComplexRes, err error) {
	return nil, service.SystemConfig().SetUserPasswordComplex(ctx, req.PasswordComplex)
}
