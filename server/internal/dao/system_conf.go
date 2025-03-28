// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalSystemConfDao is internal type for wrapping internal DAO implements.
type internalSystemConfDao = *internal.SystemConfDao

// systemConfDao is the data access object for table system_conf.
// You can define custom methods on it to extend its functionality as you wish.
type systemConfDao struct {
	internalSystemConfDao
}

var (
	// SystemConf is globally public accessible object for table system_conf operations.
	SystemConf = systemConfDao{
		internal.NewSystemConfDao(),
	}
)

// Fill with you ideas below.
