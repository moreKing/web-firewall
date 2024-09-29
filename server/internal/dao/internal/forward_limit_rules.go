// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForwardLimitRulesDao is the data access object for table forward_limit_rules.
type ForwardLimitRulesDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns ForwardLimitRulesColumns // columns contains all the column names of Table for convenient usage.
}

// ForwardLimitRulesColumns defines and stores column names for table forward_limit_rules.
type ForwardLimitRulesColumns struct {
	Id        string // 主键
	Protocol  string //
	PortType  string //
	Port      string //
	Sip       string //
	Dip       string //
	Limit     string //
	Speed     string //
	Comment   string // 备注，无意义给用户看的
	Position  string // 规则位置，重启服务时按此从小到大排序
	CreatedAt string //
	DeletedAt string //
}

// forwardLimitRulesColumns holds the columns for table forward_limit_rules.
var forwardLimitRulesColumns = ForwardLimitRulesColumns{
	Id:        "id",
	Protocol:  "protocol",
	PortType:  "port_type",
	Port:      "port",
	Sip:       "sip",
	Dip:       "dip",
	Limit:     "limit",
	Speed:     "speed",
	Comment:   "comment",
	Position:  "position",
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
}

// NewForwardLimitRulesDao creates and returns a new DAO object for table data access.
func NewForwardLimitRulesDao() *ForwardLimitRulesDao {
	return &ForwardLimitRulesDao{
		group:   "default",
		table:   "forward_limit_rules",
		columns: forwardLimitRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ForwardLimitRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ForwardLimitRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ForwardLimitRulesDao) Columns() ForwardLimitRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ForwardLimitRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ForwardLimitRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ForwardLimitRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
