// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogSettingsDao is the data access object for table log_settings.
type LogSettingsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns LogSettingsColumns // columns contains all the column names of Table for convenient usage.
}

// LogSettingsColumns defines and stores column names for table log_settings.
type LogSettingsColumns struct {
	Id            string // 主键
	Name          string // 接口名称
	Loginname     string // 登陆名
	Username      string // 用户名
	ClientIp      string // 登陆IP
	UserId        string // 用户id
	Success       string // 操作成功为true，当响应码与code==0 时才为成功
	DepartmentId  string // 部门id
	RequestMethod string // 请求方式
	RequestPath   string //
	RequestBody   string // 请求内容
	ResponseCode  string // 相应码
	ResponseError string // 错误内容
	ResponseBody  string // 相应内容
	CreatedAt     string //
}

// logSettingsColumns holds the columns for table log_settings.
var logSettingsColumns = LogSettingsColumns{
	Id:            "id",
	Name:          "name",
	Loginname:     "loginname",
	Username:      "username",
	ClientIp:      "client_ip",
	UserId:        "user_id",
	Success:       "success",
	DepartmentId:  "department_id",
	RequestMethod: "request_method",
	RequestPath:   "request_path",
	RequestBody:   "request_body",
	ResponseCode:  "response_code",
	ResponseError: "response_error",
	ResponseBody:  "response_body",
	CreatedAt:     "created_at",
}

// NewLogSettingsDao creates and returns a new DAO object for table data access.
func NewLogSettingsDao() *LogSettingsDao {
	return &LogSettingsDao{
		group:   "default",
		table:   "log_settings",
		columns: logSettingsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LogSettingsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LogSettingsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LogSettingsDao) Columns() LogSettingsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LogSettingsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LogSettingsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LogSettingsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
