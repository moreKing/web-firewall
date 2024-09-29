// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DnatRules is the golang structure of table dnat_rules for DAO operations like Where/Data.
type DnatRules struct {
	g.Meta    `orm:"table:dnat_rules, do:true"`
	Id        interface{} // 主键
	Protocol  interface{} //
	Dip       interface{} //
	Iif       interface{} //
	Port      interface{} //
	Dnat      interface{} //
	Comment   interface{} // 备注，无意义给用户看的
	Position  interface{} // 规则位置，重启服务时按此从小到大排序
	CreatedAt interface{} //
	DeletedAt interface{} //
}
