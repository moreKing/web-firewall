package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetOutputPolicyReq struct {
	g.Meta `path:"/output"  tags:"本地策略" method:"get" summary:"获取出站策略"`
}
type GetOutputPolicyRes struct {
	Data      []entity.OutputRules `json:"data"`
	Network   []model.Network      `json:"network"`
	Total     int                  `json:"total"`
	Timestamp int64                `json:"timestamp"`
}

type AddOutputPolicyReq struct {
	g.Meta `path:"/output"  tags:"本地策略" method:"post" summary:"添加出站策略"`

	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Icmp string `json:"icmp"`
	Ct   string `json:"ct"`

	Comment  string `json:"comment"`
	Add      bool   `json:"add"`      // 使用add 还是insert
	Position int    `json:"position"` // 添加规则时插入的位置
	Policy   string `json:"policy"  v:"required"`
}
type AddOutputPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceOutputPolicyReq struct {
	g.Meta   `path:"/output/:id"  tags:"本地策略" method:"put" summary:"修改出站策略"`
	ID       int64  `json:"id" path:"id" example:"1"`
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Icmp string `json:"icmp"`
	Ct   string `json:"ct"`

	Comment string `json:"comment"`

	Policy string `json:"policy"  v:"required"`
}
type ReplaceOutputPolicyRes struct {
}

type ChangeOutputPolicyPositionReq struct {
	g.Meta   `path:"/output/position/:id"  tags:"本地策略" method:"put" summary:"修改出站策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeOutputPolicyPositionRes struct {
}

type DeleteOutputPolicyReq struct {
	g.Meta `path:"/output/:id"  tags:"本地策略" method:"delete" summary:"删除出站策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteOutputPolicyRes struct {
}
