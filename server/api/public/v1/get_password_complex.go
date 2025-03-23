package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetPasswordComplexReq struct {
	g.Meta `path:"/password-complex"  tags:"公共接口" method:"get" summary:"获取密码复杂度"`
}

type GetPasswordComplexRes struct {
	*model.PasswordComplex
}
