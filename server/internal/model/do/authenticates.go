// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Authenticates is the golang structure of table authenticates for DAO operations like Where/Data.
type Authenticates struct {
	g.Meta    `orm:"table:authenticates, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 认证名称
	Mix       interface{} // 是否双因素认证，TRUE是双因素
	First     interface{} // 一步认证
	Second    interface{} // 二步认证 3.totp 4.邮件 5.短信
	CreatedAt interface{} //
	UpdatedAt interface{} //
	DeletedAt interface{} //
}
