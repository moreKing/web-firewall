package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"登录" method:"post" summary:"登出" `
}
type LogoutRes struct {
}
