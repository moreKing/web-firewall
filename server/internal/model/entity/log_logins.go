// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LogLogins is the golang structure for table log_logins.
type LogLogins struct {
	Uuid         string `json:"uuid"         orm:"uuid"          ` // 主键
	Loginname    string `json:"loginname"    orm:"loginname"     ` // 登陆名
	Username     string `json:"username"     orm:"username"      ` // 用户名
	ClientIp     string `json:"clientIp"     orm:"client_ip"     ` // 登陆IP
	UserId       int64  `json:"userId"       orm:"user_id"       ` // 用户id，登陆失败为0
	TotpCode     string `json:"totpCode"     orm:"totp_code"     ` // 手机令牌totp，防止短时间再次使用
	Success      bool   `json:"success"      orm:"success"       ` // TRUE 为登陆成功
	Online       bool   `json:"online"       orm:"online"        ` // TRUE 用户在线
	DepartmentId int64  `json:"departmentId" orm:"department_id" ` // 登陆用户所属部门，审计管理员只能查看本部门的，0所有部门都可以查看
	Log          string `json:"log"          orm:"log"           ` // 登出日志，如果登陆失败则为登陆失败日志
	LoginAt      int64  `json:"loginAt"      orm:"login_at"      ` //
	LogoutAt     int64  `json:"logoutAt"     orm:"logout_at"     ` //
}
