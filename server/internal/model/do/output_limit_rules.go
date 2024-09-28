// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OutputLimitRules is the golang structure of table output_limit_rules for DAO operations like Where/Data.
type OutputLimitRules struct {
	g.Meta    `orm:"table:output_limit_rules, do:true"`
	Id        interface{} // 主键
	Protocol  interface{} //
	Port      interface{} //
	Ip        interface{} // 指定IP，空表示所有IP
	Comment   interface{} // 备注，无意义给用户看的
	Limit     interface{} //
	Speed     interface{} //
	Position  interface{} // 规则位置，重启服务时按此从小到大排序
	CreatedAt interface{} //
	DeletedAt interface{} //
}
