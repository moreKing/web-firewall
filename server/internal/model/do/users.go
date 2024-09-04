// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta         `orm:"table:users, do:true"`
	Id             interface{} // 主键
	Loginname      interface{} // 登录名称
	Username       interface{} // 用户名
	State          interface{} // 1.启用 2.有效期 3.禁用
	Slat           interface{} // 密码加盐
	Password       interface{} // SM3加密密码
	TotpState      interface{} // 手机令牌绑定状态，true代表需要重新绑定
	TotpToken      interface{} // totp自动生成的token
	Email          interface{} //
	Mobile         interface{} //
	AuthenticateId interface{} //
	RoleId         interface{} //
	PwdUpdateAt    interface{} // 最后改密时间
	LastloginAt    interface{} // 最后登录时间
	CreatedAt      interface{} //
	UpdatedAt      interface{} //
	DeletedAt      interface{} //
}
