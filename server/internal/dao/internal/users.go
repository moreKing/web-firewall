// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	Id             string // 主键
	Loginname      string // 登录名称
	Username       string // 用户名
	State          string // 1.启用 2.有效期 3.禁用
	Slat           string // 密码加盐
	Password       string // SM3加密密码
	TotpState      string // 手机令牌绑定状态，true代表需要重新绑定
	TotpToken      string // totp自动生成的token
	Email          string //
	Mobile         string //
	AuthenticateId string //
	RoleId         string //
	PwdUpdateAt    string // 最后改密时间
	LastloginAt    string // 最后登录时间
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	Id:             "id",
	Loginname:      "loginname",
	Username:       "username",
	State:          "state",
	Slat:           "slat",
	Password:       "password",
	TotpState:      "totp_state",
	TotpToken:      "totp_token",
	Email:          "email",
	Mobile:         "mobile",
	AuthenticateId: "authenticate_id",
	RoleId:         "role_id",
	PwdUpdateAt:    "pwd_update_at",
	LastloginAt:    "lastlogin_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
