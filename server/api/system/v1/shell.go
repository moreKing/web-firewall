package v1

import "github.com/gogf/gf/v2/frame/g"

type ShellReq struct {
	g.Meta        `path:"/shell"  tags:"系统功能" method:"get" summary:"shell"`
	Authorization string `json:"Authorization" v:"required"`
}
type ShellRes struct {
}
