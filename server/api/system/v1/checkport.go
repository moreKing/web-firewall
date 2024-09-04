package v1

import "github.com/gogf/gf/v2/frame/g"

type CheckPortReq struct {
	g.Meta `path:"/check-port"  tags:"系统功能" method:"post" summary:"端口测试"`
	Ip     string `json:"ip"  v:"required"`
	Port   int    `json:"port" v:"required|min:1|max:65535"`
}
type CheckPortRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
