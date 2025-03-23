package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type RulePort struct {
	Protocol string `json:"protocol" v:"required|in:tcp,udp,tcp+udp" `
	Key      string `json:"key" v:"required"`
	Value    string `json:"value" v:"required"`
}

type DnatPort struct {
	Protocol string `json:"protocol" v:"required|in:tcp,udp,tcp+udp" `
	Key      any    `json:"key"`
	Value    any    `json:"value"`
}

type DnatRulesets struct {
	g.Meta `orm:"table:dnat_rules"`
	*entity.DnatRules
	Port []DnatPort `json:"port"`
}
