// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForwardRulesDao is the data access object for table forward_rules.
type ForwardRulesDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ForwardRulesColumns // columns contains all the column names of Table for convenient usage.
}

// ForwardRulesColumns defines and stores column names for table forward_rules.
type ForwardRulesColumns struct {
	Id        string // 主键
	Protocol  string //
	Sip       string //
	Dip       string //
	Port      string //
	Comment   string // 备注，无意义给用户看的
	Policy    string //
	Position  string // 规则位置，重启服务时按此从小到大排序
	CreatedAt string //
	DeletedAt string //
}

// forwardRulesColumns holds the columns for table forward_rules.
var forwardRulesColumns = ForwardRulesColumns{
	Id:        "id",
	Protocol:  "protocol",
	Sip:       "sip",
	Dip:       "dip",
	Port:      "port",
	Comment:   "comment",
	Policy:    "policy",
	Position:  "position",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewForwardRulesDao creates and returns a new DAO object for table data access.
func NewForwardRulesDao() *ForwardRulesDao {
	return &ForwardRulesDao{
		group:   "default",
		table:   "forward_rules",
		columns: forwardRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ForwardRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ForwardRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ForwardRulesDao) Columns() ForwardRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ForwardRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ForwardRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ForwardRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
