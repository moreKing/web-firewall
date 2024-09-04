// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SystemConf is the golang structure of table system_conf for DAO operations like Where/Data.
type SystemConf struct {
	g.Meta `orm:"table:system_conf, do:true"`
	Id     interface{} // 主键
	Name   interface{} // 配置项名称
	Config interface{} // 配置项内容
}
