// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DnatRulesDao is the data access object for table dnat_rules.
type DnatRulesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns DnatRulesColumns // columns contains all the column names of Table for convenient usage.
}

// DnatRulesColumns defines and stores column names for table dnat_rules.
type DnatRulesColumns struct {
	Id        string // 主键
	Protocol  string //
	Dip       string //
	Iif       string //
	Port      string //
	Dnat      string //
	Comment   string // 备注，无意义给用户看的
	Position  string // 规则位置，重启服务时按此从小到大排序
	CreatedAt string //
	DeletedAt string //
}

// dnatRulesColumns holds the columns for table dnat_rules.
var dnatRulesColumns = DnatRulesColumns{
	Id:        "id",
	Protocol:  "protocol",
	Dip:       "dip",
	Iif:       "iif",
	Port:      "port",
	Dnat:      "dnat",
	Comment:   "comment",
	Position:  "position",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewDnatRulesDao creates and returns a new DAO object for table data access.
func NewDnatRulesDao() *DnatRulesDao {
	return &DnatRulesDao{
		group:   "default",
		table:   "dnat_rules",
		columns: dnatRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DnatRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DnatRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DnatRulesDao) Columns() DnatRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DnatRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DnatRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DnatRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
