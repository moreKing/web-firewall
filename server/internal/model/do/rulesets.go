// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Rulesets is the golang structure of table rulesets for DAO operations like Where/Data.
type Rulesets struct {
	g.Meta    `orm:"table:rulesets, do:true"`
	Id        interface{} // 主键
	Comment   interface{} // 备注，无意义给用户看的
	Chain     interface{} // 属于链 1 入站策略 2 出站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单
	Position  interface{} // 规则位置，重启服务时按此从小到大排序
	Expr      interface{} //
	CreatedAt interface{} //
	DeletedAt interface{} //
}
