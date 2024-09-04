package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
}

type ContextUser struct {
	ID             int64  `json:"id"` // User ID.
	Authentication string // User passport.
	LoginName      string
	Password       string // 某些场景将当前登录账号密码传入到其他的地方，密码在数据库中不可逆
	ClientIP       string `json:"clientIP"`
	CreatedAt      int64  `json:"createdAt"`
	ExpiredAt      int64  `json:"expiredAt"`
}
