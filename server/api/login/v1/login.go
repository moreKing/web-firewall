package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"登录" method:"post" summary:"登录"`
	Username string `v:"required" json:"username" dc:"登录名" example:"admin"`
	Password string `v:"required" json:"password" dc:"密码" example:"admin"`
	Token    string `json:"token" dc:"邮件、短信二次认证，需要携带首次认证返回的token" example:""`
	Code     string `json:"code" dc:"动态令牌" example:""`
}
type LoginRes struct {
	Authentication string `json:"authentication,omitempty" dc:"令牌"`
	Issuer         string `json:"issuer,omitempty"`
}
