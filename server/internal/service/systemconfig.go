// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
)

type (
	ISystemConfig interface {
		GetEmail() *model.Email
		SetEmail(email *model.Email) error
		SendEmailText(addr, title, content string) error
		SendEmailStream(email *model.Email) error
		SendEmailHtml(addr, title, content string) error
		SendEmailHtmlZip(addr, title, content, fileName string, f []byte) error
		GetUserPasswordComplex() *model.PasswordComplex
		SetUserPasswordComplex(ctx context.Context, conf *model.PasswordComplex) error
		GetWebSession() *model.WebSession
		SetWebSession(ws *model.WebSession) error
		GetMessage() *model.Message
		SetMessage(messageConf *model.Message) error
		SendMessage(ctx context.Context, code, to string) error
		GetAuthConf() *model.AuthenticateConf
		SetAuthConf(ctx context.Context, a *model.AuthenticateConf) error
		GetAccountExceptionRule() *model.AccountExceptionRule
		SetAccountExceptionRule(ctx context.Context, rule *model.AccountExceptionRule) error
	}
)

var (
	localSystemConfig ISystemConfig
)

func SystemConfig() ISystemConfig {
	if localSystemConfig == nil {
		panic("implement not found for interface ISystemConfig, forgot register?")
	}
	return localSystemConfig
}

func RegisterSystemConfig(i ISystemConfig) {
	localSystemConfig = i
}
