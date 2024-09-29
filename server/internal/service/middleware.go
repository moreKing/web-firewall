// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Auth  校验登录状态是否可用
		Auth(r *ghttp.Request)
		ConfigLog(r *ghttp.Request)
		// Ctx 将自定义业务上下文变量注入到当前请求的上下文中
		Ctx(r *ghttp.Request)
		// CORS allows Cross-origin resource sharing. 提供前端跨域（不开启，前端自行代理解决）
		CORS(r *ghttp.Request)
		IsForward(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
