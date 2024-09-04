// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SqliteSequenceDao is the data access object for table sqlite_sequence.
type SqliteSequenceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SqliteSequenceColumns // columns contains all the column names of Table for convenient usage.
}

// SqliteSequenceColumns defines and stores column names for table sqlite_sequence.
type SqliteSequenceColumns struct {
	Name string //
	Seq  string //
}

// sqliteSequenceColumns holds the columns for table sqlite_sequence.
var sqliteSequenceColumns = SqliteSequenceColumns{
	Name: "name",
	Seq:  "seq",
}

// NewSqliteSequenceDao creates and returns a new DAO object for table data access.
func NewSqliteSequenceDao() *SqliteSequenceDao {
	return &SqliteSequenceDao{
		group:   "default",
		table:   "sqlite_sequence",
		columns: sqliteSequenceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SqliteSequenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SqliteSequenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SqliteSequenceDao) Columns() SqliteSequenceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SqliteSequenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SqliteSequenceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SqliteSequenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
