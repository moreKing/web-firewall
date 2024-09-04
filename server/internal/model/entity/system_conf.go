// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemConf is the golang structure for table system_conf.
type SystemConf struct {
	Id     int    `json:"id"     orm:"id"     ` // 主键
	Name   string `json:"name"   orm:"name"   ` // 配置项名称
	Config string `json:"config" orm:"config" ` // 配置项内容
}
