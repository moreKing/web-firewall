// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LogLogins is the golang structure of table log_logins for DAO operations like Where/Data.
type LogLogins struct {
	g.Meta       `orm:"table:log_logins, do:true"`
	Uuid         interface{} // 主键
	Loginname    interface{} // 登陆名
	Username     interface{} // 用户名
	ClientIp     interface{} // 登陆IP
	UserId       interface{} // 用户id，登陆失败为0
	TotpCode     interface{} // 手机令牌totp，防止短时间再次使用
	Success      interface{} // TRUE 为登陆成功
	Online       interface{} // TRUE 用户在线
	DepartmentId interface{} // 登陆用户所属部门，审计管理员只能查看本部门的，0所有部门都可以查看
	Log          interface{} // 登出日志，如果登陆失败则为登陆失败日志
	LoginAt      interface{} //
	LogoutAt     interface{} //
}
