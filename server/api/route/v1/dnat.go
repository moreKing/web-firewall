package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetDnatPolicyReq struct {
	g.Meta `path:"/dnat"  tags:"路由策略" method:"get" summary:"获取目的地址转换策略"`
}
type GetDnatPolicyRes struct {
	Data      []model.DnatRulesets `json:"data"`
	Network   []model.Network      `json:"network"`
	Total     int                  `json:"total"`
	Timestamp int64                `json:"timestamp"`
}

type AddDnatPolicyReq struct {
	g.Meta `path:"/dnat"  tags:"路由策略" method:"post" summary:"添加目的地址转换策略"`

	Dip     string           `json:"dip" v:"ipv4"`
	Iif     string           `json:"iif" v:"required"`
	Port    []model.RulePort `json:"port"`
	Dnat    string           `json:"dnat" v:"ipv4"`
	Comment string           `json:"comment"`

	Add      bool `json:"add"`      // 使用add 还是insert
	Position int  `json:"position"` // 添加规则时插入的位置

}
type AddDnatPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceDnatPolicyReq struct {
	g.Meta `path:"/dnat/:id"  tags:"路由策略" method:"put" summary:"修改目的地址转换策略"`
	ID     int64 `json:"id" path:"id" example:"1"`

	Dip     string           `json:"dip" v:"ipv4"`
	Iif     string           `json:"iif" v:"required"`
	Port    []model.RulePort `json:"port"`
	Dnat    string           `json:"dnat" v:"ipv4"`
	Comment string           `json:"comment"`
}
type ReplaceDnatPolicyRes struct {
}

type ChangeDnatPolicyPositionReq struct {
	g.Meta   `path:"/dnat/position/:id"  tags:"路由策略" method:"put" summary:"修改目的地址转换策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeDnatPolicyPositionRes struct {
}

type DeleteDnatPolicyReq struct {
	g.Meta `path:"/dnat/:id"  tags:"路由策略" method:"delete" summary:"删除目的地址转换策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteDnatPolicyRes struct {
}
