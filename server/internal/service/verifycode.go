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
	ICodeServer interface {
		VerifyCode(ctx context.Context, id int64, token, code, ip string) bool
		CreateCode(ctx context.Context, id int64, ip string, offset int) (*model.VerificationCode, error)
		VerifyTotp(ctx context.Context, code, secret string, userid int64) error
		RemoveExpireCode()
	}
)

var (
	localCodeServer ICodeServer
)

func CodeServer() ICodeServer {
	if localCodeServer == nil {
		panic("implement not found for interface ICodeServer, forgot register?")
	}
	return localCodeServer
}

func RegisterCodeServer(i ICodeServer) {
	localCodeServer = i
}
