// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
	"server/internal/model/entity"
)

type (
	ISession interface {
		// AddOnlineUser 登录成功后调用这个方法进行token注册
		AddOnlineUser(m *model.ContextUser) bool
		// RemoveOnlineUser 退出登录调用此方法取消登录状态
		RemoveOnlineUser(ctx context.Context, authentication, description string, logoutAt int64) bool
		// GetOnlineUser 校验token是否存在，存在返回对象
		GetOnlineUser(authentication string) (*model.ContextUser, bool)
		// // SetUser sets user into the session.
		//
		//	func (s *sSession) SetUser(ctx context.Context, user *entity.Users) error {
		//		return BizCtx().Get(ctx).Session.Set(consts.UserSessionKey, user)
		//	}
		//
		// GetUser retrieves and returns the user from session.
		// It returns nil if the user did not sign in.
		GetUser(ctx context.Context) *entity.Users
		// RemoveUser removes user rom session.
		RemoveUser(ctx context.Context) error
		// ClearExpiredSession 清理过期会话
		ClearExpiredSession(ctx context.Context)
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
