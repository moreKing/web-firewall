// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ForwardLimitRules is the golang structure of table forward_limit_rules for DAO operations like Where/Data.
type ForwardLimitRules struct {
	g.Meta    `orm:"table:forward_limit_rules, do:true"`
	Id        interface{} // 主键
	Protocol  interface{} //
	PortType  interface{} //
	Port      interface{} //
	Sip       interface{} //
	Dip       interface{} //
	Limit     interface{} //
	Speed     interface{} //
	Comment   interface{} // 备注，无意义给用户看的
	Position  interface{} // 规则位置，重启服务时按此从小到大排序
	CreatedAt interface{} //
	DeletedAt interface{} //
}
