package model

import (
	"github.com/gogf/gf/v2/frame/g"
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

type RulePort struct {
	Key   int `json:"key" v:"required"`
	Value int `json:"value" v:"required"`
}

type DnatRulesets struct {
	g.Meta `orm:"table:dnat_rules"`
	*entity.DnatRules
	Port []RulePort `json:"port"`
}
