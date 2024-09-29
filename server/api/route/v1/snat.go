package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetSnatPolicyReq struct {
	g.Meta `path:"/snat"  tags:"路由策略" method:"get" summary:"获取源地址转换策略"`
}
type GetSnatPolicyRes struct {
	Data      []entity.SnatRules `json:"data"`
	Network   []model.Network    `json:"network"`
	Total     int                `json:"total"`
	Timestamp int64              `json:"timestamp"`
}

type AddSnatPolicyReq struct {
	g.Meta  `path:"/snat"  tags:"路由策略" method:"post" summary:"添加源地址转换策略"`
	Sip     string `json:"sip"`
	Dip     string `json:"dip"`
	Oif     string `json:"oif"  v:"required"`
	Snat    string `json:"snat"`
	Comment string `json:"comment"`

	Add      bool `json:"add"`      // 使用add 还是insert
	Position int  `json:"position"` // 添加规则时插入的位置

}
type AddSnatPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceSnatPolicyReq struct {
	g.Meta  `path:"/snat/:id"  tags:"路由策略" method:"put" summary:"修改源地址转换策略"`
	ID      int64  `json:"id" path:"id" example:"1"`
	Sip     string `json:"sip"`
	Dip     string `json:"dip"`
	Oif     string `json:"oif" v:"required"`
	Snat    string `json:"snat"`
	Comment string `json:"comment"`
}
type ReplaceSnatPolicyRes struct {
}

type ChangeSnatPolicyPositionReq struct {
	g.Meta   `path:"/snat/position/:id"  tags:"路由策略" method:"put" summary:"修改源地址转换策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeSnatPolicyPositionRes struct {
}

type DeleteSnatPolicyReq struct {
	g.Meta `path:"/snat/:id"  tags:"路由策略" method:"delete" summary:"删除源地址转换策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteSnatPolicyRes struct {
}
