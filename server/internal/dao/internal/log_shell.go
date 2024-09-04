// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogShellDao is the data access object for table log_shell.
type LogShellDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LogShellColumns // columns contains all the column names of Table for convenient usage.
}

// LogShellColumns defines and stores column names for table log_shell.
type LogShellColumns struct {
	Id        string // 主键
	Loginname string // 登陆名
	Username  string // 用户名
	ClientIp  string // 登陆IP
	UserId    string // 用户id
	Success   string // 操作成功为true
	Online    string // true在线
	Filename  string //
	Md5       string // 记录完成后防止记录被篡改，需要记录文件md5值
	Size      string //
	LogoutAt  string //
	CreatedAt string //
}

// logShellColumns holds the columns for table log_shell.
var logShellColumns = LogShellColumns{
	Id:        "id",
	Loginname: "loginname",
	Username:  "username",
	ClientIp:  "client_ip",
	UserId:    "user_id",
	Success:   "success",
	Online:    "online",
	Filename:  "filename",
	Md5:       "md5",
	Size:      "size",
	LogoutAt:  "logout_at",
	CreatedAt: "created_at",
}

// NewLogShellDao creates and returns a new DAO object for table data access.
func NewLogShellDao() *LogShellDao {
	return &LogShellDao{
		group:   "default",
		table:   "log_shell",
		columns: logShellColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LogShellDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LogShellDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LogShellDao) Columns() LogShellColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LogShellDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LogShellDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LogShellDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
