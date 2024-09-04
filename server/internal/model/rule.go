package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/model/entity"
	"server/utility/nftables"
)

type Rulesets struct {
	gmeta.Meta `orm:"table:rulesets"`
	*entity.Rulesets
	Handle string                `json:"handle"`
	Expr   []nftables.Expression `json:"expr"`
}
