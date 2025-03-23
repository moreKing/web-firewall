// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthenticatesDao is the data access object for table authenticates.
type AuthenticatesDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AuthenticatesColumns // columns contains all the column names of Table for convenient usage.
}

// AuthenticatesColumns defines and stores column names for table authenticates.
type AuthenticatesColumns struct {
	Id        string // 主键
	Name      string // 认证名称
	Mix       string // 是否双因素认证，TRUE是双因素
	First     string // 一步认证
	Second    string // 二步认证 3.totp 4.邮件 5.短信
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// authenticatesColumns holds the columns for table authenticates.
var authenticatesColumns = AuthenticatesColumns{
	Id:        "id",
	Name:      "name",
	Mix:       "mix",
	First:     "first",
	Second:    "second",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAuthenticatesDao creates and returns a new DAO object for table data access.
func NewAuthenticatesDao() *AuthenticatesDao {
	return &AuthenticatesDao{
		group:   "default",
		table:   "authenticates",
		columns: authenticatesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuthenticatesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuthenticatesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuthenticatesDao) Columns() AuthenticatesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuthenticatesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuthenticatesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuthenticatesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
