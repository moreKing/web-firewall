package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SetUserPasswordReq struct {
	g.Meta      `path:"/set-password" tags:"公共接口" method:"put" summary:"更新当前登录用户密码"`
	OldPassword string `json:"oldPassword" dc:"旧密码" v:"required#请填写旧密码"`
	NewPassword string `json:"newPassword" dc:"新密码" v:"required|not-eq:OldPassword#请填写新密码|新密码不能与旧密码一致"`
}
type SetUserPasswordRes struct {
}
