// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// InputLimitRules is the golang structure for table input_limit_rules.
type InputLimitRules struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Protocol  string `json:"protocol"  orm:"protocol"   ` //
	Port      string `json:"port"      orm:"port"       ` //
	Ip        string `json:"ip"        orm:"ip"         ` // 指定IP，空表示所有IP
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Limit     int    `json:"limit"     orm:"limit"      ` //
	Speed     string `json:"speed"     orm:"speed"      ` //
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
