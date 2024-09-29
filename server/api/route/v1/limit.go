package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetLimitPolicyReq struct {
	g.Meta `path:"/limit"  tags:"路由策略" method:"get" summary:"获取转发流控策略"`
}
type GetLimitPolicyRes struct {
	Data      []entity.ForwardLimitRules `json:"data"`
	Network   []model.Network            `json:"network"`
	Total     int                        `json:"total"`
	Timestamp int64                      `json:"timestamp"`
}

type AddLimitPolicyReq struct {
	g.Meta `path:"/limit"  tags:"路由策略" method:"post" summary:"添加转发流控策略"`

	Protocol string `json:"protocol"`
	PortType string `json:"portType" v:"in:sport,dport"`
	Port     string `json:"port"`
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Limit    int    `json:"limit" v:"required"`
	Speed    string `json:"speed" v:"required|in:kb/s,mb/s,kb/m,mb/m"`
	Comment  string `json:"comment"`

	Add      bool `json:"add"`      // 使用add 还是insert
	Position int  `json:"position"` // 添加规则时插入的位置

}
type AddLimitPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceLimitPolicyReq struct {
	g.Meta `path:"/limit/:id"  tags:"路由策略" method:"put" summary:"修改转发流控策略"`
	ID     int64 `json:"id" path:"id" example:"1"`

	Protocol string `json:"protocol"`
	PortType string `json:"portType"  v:"in:sport,dport"`
	Port     string `json:"port"`
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Limit    int    `json:"limit" v:"required"`
	Speed    string `json:"speed" v:"required|in:kb/s,mb/s,kb/m,mb/m"`
	Comment  string `json:"comment"`
}
type ReplaceLimitPolicyRes struct {
}

type ChangeLimitPolicyPositionReq struct {
	g.Meta   `path:"/limit/position/:id"  tags:"路由策略" method:"put" summary:"修改转发流控策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeLimitPolicyPositionRes struct {
}

type DeleteLimitPolicyReq struct {
	g.Meta `path:"/limit/:id"  tags:"路由策略" method:"delete" summary:"删除转发流控策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteLimitPolicyRes struct {
}
