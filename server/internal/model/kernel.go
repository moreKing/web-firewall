package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/model/entity"
)

type Kernel struct {
	Forward bool `json:"forward" v:"required"` // false 关闭 路径 /proc/sys/net/ipv4/ip_forward
}
type SystemConfKernel struct {
	gmeta.Meta `orm:"table:system_conf"`
	*entity.SystemConf
	Config Kernel
}
