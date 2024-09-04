package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SetProfileReq struct {
	g.Meta         `path:"/profile"  tags:"公共接口" method:"put" summary:"更新当前登录用户信息"`
	Username       string `json:"username" v:"required"`
	Email          string `json:"email"  v:"email"`
	Phone          string `json:"phone"  v:"phone-loose"`
	AuthenticateId int    `json:"authenticateId" v:"min:0|max:3"`
	Bind           bool   `json:"bind" dc:"表示手机令牌作为二步认证 是否重新绑定"`
}
type SetProfileRes struct {
}
