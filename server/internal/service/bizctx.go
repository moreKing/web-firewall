// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IBizCtx interface {
		// Init initializes and injects custom business context object into request context.
		Init(r *ghttp.Request, customCtx *model.Context)
		// Get retrieves and returns the user object from context.
		// It returns nil if nothing found in given context.
		Get(ctx context.Context) *model.Context
		// SetUser injects business user object into context.
		SetUser(ctx context.Context, ctxUser *model.ContextUser)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
