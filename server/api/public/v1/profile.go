package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type ProfileReq struct {
	g.Meta `path:"/profile"  tags:"公共接口" method:"get" summary:"获取当前登录用户信息"`
}
type ProfileRes struct {
	UserInfo    *model.UserInfo `json:"user"`
	Routes      []string        `json:"routes"` // 用户页面访问权限
	Home        string          `json:"home"`
	Permissions *[]string       `json:"permissions"` //用户真正的接口访问权限
	Email       bool            `json:"email"`
	Message     bool            `json:"message"`
}
