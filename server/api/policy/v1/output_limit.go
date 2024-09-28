package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/internal/model/entity"
)

type GetOutputLimitPolicyReq struct {
	g.Meta `path:"/output-limit"  tags:"策略管理" method:"get" summary:"获取出站流控策略"`
}
type GetOutputLimitPolicyRes struct {
	Data      []entity.OutputLimitRules `json:"data"`
	Network   []model.Network           `json:"network"`
	Total     int                       `json:"total"`
	Timestamp int64                     `json:"timestamp"`
}

type AddOutputLimitPolicyReq struct {
	g.Meta `path:"/output-limit"  tags:"策略管理" method:"post" summary:"添加出站流控策略"`

	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Limit    int    `json:"limit" v:"required"`
	Speed    string `json:"speed" v:"in:kb/s,mb/s,kb/m,mb/m"`

	Comment  string `json:"comment"`
	Add      bool   `json:"add"`      // 使用add 还是insert
	Position int    `json:"position"` // 添加规则时插入的位置
}
type AddOutputLimitPolicyRes struct {
	Id int64 `json:"id"`
}

type ReplaceOutputLimitPolicyReq struct {
	g.Meta   `path:"/output-limit/:id"  tags:"策略管理" method:"put" summary:"修改出站流控策略"`
	ID       int64  `json:"id" path:"id" example:"1"`
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`

	Limit int    `json:"limit" v:"required"`
	Speed string `json:"speed" v:"in:kb/s,mb/s,kb/m,mb/m"`

	Comment string `json:"comment"`
}
type ReplaceOutputLimitPolicyRes struct {
}

type ChangeOutputLimitPolicyPositionReq struct {
	g.Meta   `path:"/output-limit/position/:id"  tags:"策略管理" method:"put" summary:"修改出站流控策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangeOutputLimitPolicyPositionRes struct {
}

type DeleteOutputLimitPolicyReq struct {
	g.Meta `path:"/output-limit/:id"  tags:"策略管理" method:"delete" summary:"删除出站流控策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeleteOutputLimitPolicyRes struct {
}
