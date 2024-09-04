// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model/entity"
)

type (
	IUser interface {
		// SetNativePassword 设置用户本地登录密码
		SetNativePassword(ctx context.Context, userId int64, newPassword string) (err error)
		// GetUserByID 查询指定用户信息
		GetUserByID(ctx context.Context, id int64) (user *entity.Users, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
