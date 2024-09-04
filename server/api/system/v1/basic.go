package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

type GetWebConfigReq struct {
	g.Meta `path:"/session-config"  tags:"基本设置" method:"get" summary:"获取会话配置"`
}
type GetWebConfigRes struct {
	*model.WebSession
}

type SetWebConfigReq struct {
	g.Meta `path:"/session-config"  tags:"基本设置" method:"put" summary:"设置会话配置"`
	*model.WebSession
}
type SetWebConfigRes struct {
}

// password-complex
type SetPasswordComplexReq struct {
	g.Meta `path:"/password-complex"  tags:"基本设置" method:"put" summary:"设置本地密码规则"`
	*model.PasswordComplex
}
type SetPasswordComplexRes struct {
}

// email
type GetEmailReq struct {
	g.Meta `path:"/email"  tags:"基本设置" method:"get" summary:"获取邮箱配置"`
}
type GetEmailRes struct {
	*model.Email
}

type SetEmailReq struct {
	g.Meta `path:"/email"  tags:"基本设置" method:"put" summary:"修改邮箱配置"`
	*model.Email
}
type SetEmailRes struct {
	*model.Email
}

type TestEmailReq struct {
	g.Meta `path:"/email"  tags:"基本设置" method:"post" summary:"邮件测试"`
	To     string `json:"to" v:"required|email"`
}
type TestEmailRes struct {
}

// 短信
type GetMessageReq struct {
	g.Meta `path:"/message"  tags:"基本设置" method:"get" summary:"获取短信配置"`
}
type GetMessageRes struct {
	*model.Message
}

type SetMessageReq struct {
	g.Meta `path:"/message"  tags:"基本设置" method:"put" summary:"修改短信配置"`
	*model.Message
}
type SetMessageRes struct {
	*model.Message
}

type TestMessageReq struct {
	g.Meta `path:"/message"  tags:"基本设置" method:"post" summary:"短信测试"`
	To     string `json:"to" v:"required|phone-loose"`
}
type TestMessageRes struct {
}

// GetAuthConfReq 认证配置
type GetAuthConfReq struct {
	g.Meta `path:"/auth-conf"  tags:"基本设置" method:"get" summary:"获取登录配置"`
}
type GetAuthConfRes struct {
	*model.AuthenticateConf
}

type SetAuthConfReq struct {
	g.Meta `path:"/auth-conf"  tags:"基本设置" method:"put" summary:"修改登录配置"`
	*model.AuthenticateConf
}
type SetAuthConfRes struct {
}
