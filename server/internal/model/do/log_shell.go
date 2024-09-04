// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LogShell is the golang structure of table log_shell for DAO operations like Where/Data.
type LogShell struct {
	g.Meta    `orm:"table:log_shell, do:true"`
	Id        interface{} // 主键
	Loginname interface{} // 登陆名
	Username  interface{} // 用户名
	ClientIp  interface{} // 登陆IP
	UserId    interface{} // 用户id
	Success   interface{} // 操作成功为true
	Online    interface{} // true在线
	Filename  interface{} //
	Md5       interface{} // 记录完成后防止记录被篡改，需要记录文件md5值
	Size      interface{} //
	LogoutAt  interface{} //
	CreatedAt interface{} //
}
