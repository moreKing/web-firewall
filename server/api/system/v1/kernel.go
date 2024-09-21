package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetKernelReq struct {
	g.Meta `path:"/kernel"  tags:"内核参数" method:"get" summary:"获取内核参数"`
}
type GetKernelRes struct {
	*model.Kernel
}

type SetKernelReq struct {
	g.Meta `path:"/kernel"  tags:"内核参数" method:"put" summary:"修改内核参数"`
	*model.Kernel
}
type SetKernelRes struct {
}
