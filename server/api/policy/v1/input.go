package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetInputPolicyReq struct {
	g.Meta `path:"/input"  tags:"本地策略" method:"get" summary:"获取入站策略"`
}
type GetInputPolicyRes struct {
	Data      []entity.InputRules `json:"data"`
	Network   []model.Network     `json:"network"`
	Total     int                 `json:"total"`
	Timestamp int64               `json:"timestamp"`
}

type AddInputPolicyReq struct {
	g.Meta `path:"/input"  tags:"本地策略" method:"post" summary:"添加入站策略"`

	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Icmp string `json:"icmp"`
	Ct   string `json:"ct"`

	Comment  string `json:"comment"`
	Add      bool   `json:"add"`      // 使用add 还是insert
	Position int    `json:"position"` // 添加规则时插入的位置
	Policy   string `json:"policy" v:"required"`
}
type AddInputPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceInputPolicyReq struct {
	g.Meta   `path:"/input/:id"  tags:"本地策略" method:"put" summary:"修改入站策略"`
	ID       int64  `json:"id" path:"id" example:"1"`
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Icmp string `json:"icmp"`
	Ct   string `json:"ct"`

	Comment string `json:"comment"`

	Policy string `json:"policy" v:"required"`
}
type ReplaceInputPolicyRes struct {
}

type ChangeInputPolicyPositionReq struct {
	g.Meta   `path:"/input/position/:id"  tags:"本地策略" method:"put" summary:"修改入站策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeInputPolicyPositionRes struct {
}

type DeleteInputPolicyReq struct {
	g.Meta `path:"/input/:id"  tags:"本地策略" method:"delete" summary:"删除入站策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteInputPolicyRes struct {
}
