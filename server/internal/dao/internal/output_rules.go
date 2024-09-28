// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OutputRulesDao is the data access object for table output_rules.
type OutputRulesDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns OutputRulesColumns // columns contains all the column names of Table for convenient usage.
}

// OutputRulesColumns defines and stores column names for table output_rules.
type OutputRulesColumns struct {
	Id        string // 主键
	Protocol  string //
	Port      string //
	Ip        string // 指定目的IP，空表示所有目的IP
	Ct        string //
	Icmp      string //
	Comment   string // 备注，无意义给用户看的
	Policy    string //
	Position  string // 规则位置，重启服务时按此从小到大排序
	CreatedAt string //
	DeletedAt string //
}

// outputRulesColumns holds the columns for table output_rules.
var outputRulesColumns = OutputRulesColumns{
	Id:        "id",
	Protocol:  "protocol",
	Port:      "port",
	Ip:        "ip",
	Ct:        "ct",
	Icmp:      "icmp",
	Comment:   "comment",
	Policy:    "policy",
	Position:  "position",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewOutputRulesDao creates and returns a new DAO object for table data access.
func NewOutputRulesDao() *OutputRulesDao {
	return &OutputRulesDao{
		group:   "default",
		table:   "output_rules",
		columns: outputRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OutputRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OutputRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OutputRulesDao) Columns() OutputRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OutputRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OutputRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OutputRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
