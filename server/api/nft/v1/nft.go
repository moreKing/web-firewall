package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
	"server/utility/nftables"
)

// 增删改查，不返回单个，查询直接返回某个链全部规则

type GetPolicyReq struct {
	g.Meta `path:"/policy/:chain"  tags:"策略管理" method:"get" summary:"添加策略"`
	Chain  int `json:"chain"  path:"chain"  v:"min:1|max:7" dc:"1 出站策略 2 入站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单"` //1:出站策略 2: 入站策略 3：目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单

}
type GetPolicyRes struct {
	Data      []model.Rulesets `json:"data"`
	Total     int              `json:"total"`
	Timestamp int64            `json:"timestamp"`
}

type AddPolicyReq struct {
	g.Meta   `path:"/policy/:chain"  tags:"策略管理" method:"post" summary:"添加策略"`
	Chain    int                   `json:"chain" path:"chain" v:"min:1|max:7" dc:"1 出站策略 2 入站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单"` //1:出站策略 2: 入站策略 3：目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单
	Comment  string                `json:"comment"`
	Add      bool                  `json:"add"`      // 使用add 还是insert
	Position int                   `json:"position"` // 添加规则时插入的位置
	Expr     []nftables.Expression `json:"expr"`
}
type AddPolicyRes struct {
}

type ReplacePolicyReq struct {
	g.Meta  `path:"/policy/:id"  tags:"策略管理" method:"put" summary:"修改策略"`
	ID      int64                 `json:"id" path:"id" example:"1"`
	Comment string                `json:"comment"`
	Expr    []nftables.Expression `json:"expr"`
}
type ReplacePolicyRes struct {
}

type ChangePolicyPositionReq struct {
	g.Meta   `path:"/policy/position/:id"  tags:"策略管理" method:"put" summary:"修改策略位置"`
	ID       int64 `json:"id" path:"id" example:"1"`
	Add      bool  `json:"add"`      // 使用add 还是insert
	Position int64 `json:"position"` // 添加规则时插入的位置
}
type ChangePolicyPositionRes struct {
}

type DeletePolicyReq struct {
	g.Meta `path:"/policy/:id"  tags:"策略管理" method:"delete" summary:"删除策略"`
	ID     int64 `json:"id" path:"id" example:"1"`
}
type DeletePolicyRes struct {
}
