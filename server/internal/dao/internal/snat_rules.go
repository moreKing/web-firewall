// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SnatRulesDao is the data access object for table snat_rules.
type SnatRulesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SnatRulesColumns // columns contains all the column names of Table for convenient usage.
}

// SnatRulesColumns defines and stores column names for table snat_rules.
type SnatRulesColumns struct {
	Id        string // 主键
	Sip       string //
	Dip       string //
	Oif       string //
	Snat      string //
	Comment   string // 备注，无意义给用户看的
	Position  string // 规则位置，重启服务时按此从小到大排序
	CreatedAt string //
	DeletedAt string //
}

// snatRulesColumns holds the columns for table snat_rules.
var snatRulesColumns = SnatRulesColumns{
	Id:        "id",
	Sip:       "sip",
	Dip:       "dip",
	Oif:       "oif",
	Snat:      "snat",
	Comment:   "comment",
	Position:  "position",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewSnatRulesDao creates and returns a new DAO object for table data access.
func NewSnatRulesDao() *SnatRulesDao {
	return &SnatRulesDao{
		group:   "default",
		table:   "snat_rules",
		columns: snatRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SnatRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SnatRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SnatRulesDao) Columns() SnatRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SnatRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SnatRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SnatRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
