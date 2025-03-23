package model

// VerificationCode 短信邮件验证码
type VerificationCode struct {
	Token     string
	ClientIP  string
	UserId    int64
	Code      string
	CreatedAt int64
	ExpiredAt int64
}
