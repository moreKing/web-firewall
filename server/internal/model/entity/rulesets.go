// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Rulesets is the golang structure for table rulesets.
type Rulesets struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Chain     int    `json:"chain"     orm:"chain"      ` // 属于链 1 入站策略 2 出站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	Expr      string `json:"expr"      orm:"expr"       ` //
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
