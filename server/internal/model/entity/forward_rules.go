// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ForwardRules is the golang structure for table forward_rules.
type ForwardRules struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Protocol  string `json:"protocol"  orm:"protocol"   ` //
	Sip       string `json:"sip"       orm:"sip"        ` //
	Dip       string `json:"dip"       orm:"dip"        ` //
	Port      string `json:"port"      orm:"port"       ` //
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Policy    string `json:"policy"    orm:"policy"     ` //
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
