// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ForwardRules is the golang structure of table forward_rules for DAO operations like Where/Data.
type ForwardRules struct {
	g.Meta    `orm:"table:forward_rules, do:true"`
	Id        interface{} // 主键
	Protocol  interface{} //
	Sip       interface{} //
	Dip       interface{} //
	Port      interface{} //
	Comment   interface{} // 备注，无意义给用户看的
	Policy    interface{} //
	Position  interface{} // 规则位置，重启服务时按此从小到大排序
	CreatedAt interface{} //
	DeletedAt interface{} //
}
