package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetInputLimitPolicyReq struct {
	g.Meta `path:"/input-limit"  tags:"策略管理" method:"get" summary:"获取入站流控策略"`
}
type GetInputLimitPolicyRes struct {
	Data      []entity.InputLimitRules `json:"data"`
	Network   []model.Network          `json:"network"`
	Total     int                      `json:"total"`
	Timestamp int64                    `json:"timestamp"`
}

type AddInputLimitPolicyReq struct {
	g.Meta `path:"/input-limit"  tags:"策略管理" method:"post" summary:"添加入站流控策略"`

	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Limit    int    `json:"limit" v:"required"`
	Speed    string `json:"speed" v:"in:kb/s,mb/s,kb/m,mb/m"`

	Comment  string `json:"comment"`
	Add      bool   `json:"add"`      // 使用add 还是insert
	Position int    `json:"position"` // 添加规则时插入的位置
}
type AddInputLimitPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceInputLimitPolicyReq struct {
	g.Meta   `path:"/input-limit/:id"  tags:"策略管理" method:"put" summary:"修改入站流控策略"`
	ID       int64  `json:"id" path:"id" example:"1"`
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Limit int    `json:"limit" v:"required"`
	Speed string `json:"speed" v:"in:kb/s,mb/s,kb/m,mb/m"`

	Comment string `json:"comment"`
}
type ReplaceInputLimitPolicyRes struct {
}

type ChangeInputLimitPolicyPositionReq struct {
	g.Meta   `path:"/input-limit/position/:id"  tags:"策略管理" method:"put" summary:"修改入站流控策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeInputLimitPolicyPositionRes struct {
}

type DeleteInputLimitPolicyReq struct {
	g.Meta `path:"/input-limit/:id"  tags:"策略管理" method:"delete" summary:"删除入站流控策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteInputLimitPolicyRes struct {
}
