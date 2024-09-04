// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LogShell is the golang structure for table log_shell.
type LogShell struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Loginname string `json:"loginname" orm:"loginname"  ` // 登陆名
	Username  string `json:"username"  orm:"username"   ` // 用户名
	ClientIp  string `json:"clientIp"  orm:"client_ip"  ` // 登陆IP
	UserId    int64  `json:"userId"    orm:"user_id"    ` // 用户id
	Success   bool   `json:"success"   orm:"success"    ` // 操作成功为true
	Online    bool   `json:"online"    orm:"online"     ` // true在线
	Filename  string `json:"filename"  orm:"filename"   ` //
	Md5       string `json:"md5"       orm:"md5"        ` // 记录完成后防止记录被篡改，需要记录文件md5值
	Size      int64  `json:"size"      orm:"size"       ` //
	LogoutAt  int64  `json:"logoutAt"  orm:"logout_at"  ` //
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
}
