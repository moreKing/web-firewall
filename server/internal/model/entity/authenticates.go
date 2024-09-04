// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Authenticates is the golang structure for table authenticates.
type Authenticates struct {
	Id        int64  `json:"id"        orm:"id"         ` // 主键
	Name      string `json:"name"      orm:"name"       ` // 认证名称
	Mix       bool   `json:"mix"       orm:"mix"        ` // 是否双因素认证，TRUE是双因素
	First     int    `json:"first"     orm:"first"      ` // 一步认证
	Second    int    `json:"second"    orm:"second"     ` // 二步认证 3.totp 4.邮件 5.短信
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` //
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt int64  `json:"deletedAt" orm:"deleted_at" ` //
}
