package public

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"

	"server/api/public/v1"
)

func (c *ControllerV1) GetPasswordComplex(ctx context.Context, req *v1.GetPasswordComplexReq) (res *v1.GetPasswordComplexRes, err error) {

	conf := &model.SystemConfPasswordComplex{}
	err = dao.SystemConf.Ctx(ctx).Where(dao.SystemConf.Columns().Id, 1).Scan(conf)
	if err != nil {
		return nil, err
	}
	res = &v1.GetPasswordComplexRes{PasswordComplex: service.SystemConfig().GetUserPasswordComplex()}
	return
}
