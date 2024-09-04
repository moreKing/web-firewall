// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RulesetsDao is the data access object for table rulesets.
type RulesetsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RulesetsColumns // columns contains all the column names of Table for convenient usage.
}

// RulesetsColumns defines and stores column names for table rulesets.
type RulesetsColumns struct {
	Id        string // 主键
	Comment   string // 备注，无意义给用户看的
	Chain     string // 属于链 1 入站策略 2 出站策略 3 目的地址转换 4 源地址转换 5 入站限流 6 出站限流 7 ip黑白名单
	Position  string // 规则位置，重启服务时按此从小到大排序
	Expr      string //
	CreatedAt string //
	DeletedAt string //
}

// rulesetsColumns holds the columns for table rulesets.
var rulesetsColumns = RulesetsColumns{
	Id:        "id",
	Comment:   "comment",
	Chain:     "chain",
	Position:  "position",
	Expr:      "expr",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewRulesetsDao creates and returns a new DAO object for table data access.
func NewRulesetsDao() *RulesetsDao {
	return &RulesetsDao{
		group:   "default",
		table:   "rulesets",
		columns: rulesetsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RulesetsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RulesetsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RulesetsDao) Columns() RulesetsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RulesetsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RulesetsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RulesetsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
