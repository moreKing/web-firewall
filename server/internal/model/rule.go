package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type RulePort struct {
	Key   int `json:"key" v:"required"`
	Value int `json:"value" v:"required"`
}

type DnatRulesets struct {
	g.Meta `orm:"table:dnat_rules"`
	*entity.DnatRules
	Port []RulePort `json:"port"`
}
