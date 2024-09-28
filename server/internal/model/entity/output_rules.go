// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OutputRules is the golang structure for table output_rules.
type OutputRules struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Protocol  string `json:"protocol"  orm:"protocol"   ` //
	Port      string `json:"port"      orm:"port"       ` //
	Ip        string `json:"ip"        orm:"ip"         ` // 指定目的IP，空表示所有目的IP
	Ct        string `json:"ct"        orm:"ct"         ` //
	Icmp      string `json:"icmp"      orm:"icmp"       ` //
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Policy    string `json:"policy"    orm:"policy"     ` //
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
