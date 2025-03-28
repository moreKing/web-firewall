// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalDnatRulesDao is internal type for wrapping internal DAO implements.
type internalDnatRulesDao = *internal.DnatRulesDao

// dnatRulesDao is the data access object for table dnat_rules.
// You can define custom methods on it to extend its functionality as you wish.
type dnatRulesDao struct {
	internalDnatRulesDao
}

var (
	// DnatRules is globally public accessible object for table dnat_rules operations.
	DnatRules = dnatRulesDao{
		internal.NewDnatRulesDao(),
	}
)

// Fill with you ideas below.
