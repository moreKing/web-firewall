// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LogSettings is the golang structure of table log_settings for DAO operations like Where/Data.
type LogSettings struct {
	g.Meta        `orm:"table:log_settings, do:true"`
	Id            interface{} // 主键
	Name          interface{} // 接口名称
	Loginname     interface{} // 登陆名
	Username      interface{} // 用户名
	ClientIp      interface{} // 登陆IP
	UserId        interface{} // 用户id
	Success       interface{} // 操作成功为true，当响应码与code==0 时才为成功
	DepartmentId  interface{} // 部门id
	RequestMethod interface{} // 请求方式
	RequestPath   interface{} //
	RequestBody   interface{} // 请求内容
	ResponseCode  interface{} // 相应码
	ResponseError interface{} // 错误内容
	ResponseBody  interface{} // 相应内容
	CreatedAt     interface{} //
}
