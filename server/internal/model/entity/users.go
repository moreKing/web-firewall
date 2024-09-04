// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Users is the golang structure for table users.
type Users struct {
	Id             int64  `json:"id"             orm:"id"              ` // 主键
	Loginname      string `json:"loginname"      orm:"loginname"       ` // 登录名称
	Username       string `json:"username"       orm:"username"        ` // 用户名
	State          int    `json:"state"          orm:"state"           ` // 1.启用 2.有效期 3.禁用
	Slat           string `json:"slat"           orm:"slat"            ` // 密码加盐
	Password       string `json:"password"       orm:"password"        ` // SM3加密密码
	TotpState      bool   `json:"totpState"      orm:"totp_state"      ` // 手机令牌绑定状态，true代表需要重新绑定
	TotpToken      string `json:"totpToken"      orm:"totp_token"      ` // totp自动生成的token
	Email          string `json:"email"          orm:"email"           ` //
	Mobile         string `json:"mobile"         orm:"mobile"          ` //
	AuthenticateId int64  `json:"authenticateId" orm:"authenticate_id" ` //
	RoleId         int64  `json:"roleId"         orm:"role_id"         ` //
	PwdUpdateAt    int64  `json:"pwdUpdateAt"    orm:"pwd_update_at"   ` // 最后改密时间
	LastloginAt    int64  `json:"lastloginAt"    orm:"lastlogin_at"    ` // 最后登录时间
	CreatedAt      int64  `json:"createdAt"      orm:"created_at"      ` //
	UpdatedAt      int64  `json:"updatedAt"      orm:"updated_at"      ` //
	DeletedAt      int64  `json:"deletedAt"      orm:"deleted_at"      ` //
}
