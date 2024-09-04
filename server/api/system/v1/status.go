package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetSystemStatusReq struct {
	g.Meta `path:"/status"  tags:"系统管理" method:"get" summary:"获取系统状态"`
}
type GetSystemStatusRes struct {
	*model.SystemStatus
}
