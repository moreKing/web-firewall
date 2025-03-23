package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetHomeReq struct {
	g.Meta `path:"/home"  tags:"系统管理" method:"get" summary:"home页面信息"`
}
type GetHomeRes struct {
	*model.Home
}
