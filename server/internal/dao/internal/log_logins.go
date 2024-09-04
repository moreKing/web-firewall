// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogLoginsDao is the data access object for table log_logins.
type LogLoginsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns LogLoginsColumns // columns contains all the column names of Table for convenient usage.
}

// LogLoginsColumns defines and stores column names for table log_logins.
type LogLoginsColumns struct {
	Uuid         string // 主键
	Loginname    string // 登陆名
	Username     string // 用户名
	ClientIp     string // 登陆IP
	UserId       string // 用户id，登陆失败为0
	TotpCode     string // 手机令牌totp，防止短时间再次使用
	Success      string // TRUE 为登陆成功
	Online       string // TRUE 用户在线
	DepartmentId string // 登陆用户所属部门，审计管理员只能查看本部门的，0所有部门都可以查看
	Log          string // 登出日志，如果登陆失败则为登陆失败日志
	LoginAt      string //
	LogoutAt     string //
}

// logLoginsColumns holds the columns for table log_logins.
var logLoginsColumns = LogLoginsColumns{
	Uuid:         "uuid",
	Loginname:    "loginname",
	Username:     "username",
	ClientIp:     "client_ip",
	UserId:       "user_id",
	TotpCode:     "totp_code",
	Success:      "success",
	Online:       "online",
	DepartmentId: "department_id",
	Log:          "log",
	LoginAt:      "login_at",
	LogoutAt:     "logout_at",
}

// NewLogLoginsDao creates and returns a new DAO object for table data access.
func NewLogLoginsDao() *LogLoginsDao {
	return &LogLoginsDao{
		group:   "default",
		table:   "log_logins",
		columns: logLoginsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LogLoginsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LogLoginsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LogLoginsDao) Columns() LogLoginsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LogLoginsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LogLoginsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LogLoginsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
