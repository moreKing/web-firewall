package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetForwardPolicyReq struct {
	g.Meta `path:"/forward"  tags:"路由策略" method:"get" summary:"获取转发策略"`
}
type GetForwardPolicyRes struct {
	Data      []entity.ForwardRules `json:"data"`
	Network   []model.Network       `json:"network"`
	Total     int                   `json:"total"`
	Timestamp int64                 `json:"timestamp"`
}

type AddForwardPolicyReq struct {
	g.Meta `path:"/forward"  tags:"路由策略" method:"post" summary:"添加转发策略"`

	Protocol string `json:"protocol"`
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Port     string `json:"port"`
	Policy   string `json:"policy" v:"required|in:accept,drop"`
	Comment  string `json:"comment"`

	Add      bool `json:"add"`      // 使用add 还是insert
	Position int  `json:"position"` // 添加规则时插入的位置

}
type AddForwardPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceForwardPolicyReq struct {
	g.Meta `path:"/forward/:id"  tags:"路由策略" method:"put" summary:"修改转发策略"`
	ID     int64 `json:"id" path:"id" example:"1"`

	Protocol string `json:"protocol"`
	Sip      string `json:"sip"`
	Dip      string `json:"dip"`
	Port     string `json:"port"`
	Policy   string `json:"policy" v:"required|in:accept,drop"`
	Comment  string `json:"comment"`
}
type ReplaceForwardPolicyRes struct {
}

type ChangeForwardPolicyPositionReq struct {
	g.Meta   `path:"/forward/position/:id"  tags:"路由策略" method:"put" summary:"修改转发策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeForwardPolicyPositionRes struct {
}

type DeleteForwardPolicyReq struct {
	g.Meta `path:"/forward/:id"  tags:"路由策略" method:"delete" summary:"删除转发策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteForwardPolicyRes struct {
}
