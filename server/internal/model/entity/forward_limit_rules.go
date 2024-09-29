// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ForwardLimitRules is the golang structure for table forward_limit_rules.
type ForwardLimitRules struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Protocol  string `json:"protocol"  orm:"protocol"   ` //
	PortType  string `json:"portType"  orm:"port_type"  ` //
	Port      string `json:"port"      orm:"port"       ` //
	Sip       string `json:"sip"       orm:"sip"        ` //
	Dip       string `json:"dip"       orm:"dip"        ` //
	Limit     int    `json:"limit"     orm:"limit"      ` //
	Speed     string `json:"speed"     orm:"speed"      ` //
	Comment   string `json:"comment"   orm:"comment"    ` // 备注，无意义给用户看的
	Position  int    `json:"position"  orm:"position"   ` // 规则位置，重启服务时按此从小到大排序
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
