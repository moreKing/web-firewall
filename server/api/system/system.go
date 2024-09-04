// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"server/api/system/v1"
)

type ISystemV1 interface {
	GetWebConfig(ctx context.Context, req *v1.GetWebConfigReq) (res *v1.GetWebConfigRes, err error)
	SetWebConfig(ctx context.Context, req *v1.SetWebConfigReq) (res *v1.SetWebConfigRes, err error)
	SetPasswordComplex(ctx context.Context, req *v1.SetPasswordComplexReq) (res *v1.SetPasswordComplexRes, err error)
	GetEmail(ctx context.Context, req *v1.GetEmailReq) (res *v1.GetEmailRes, err error)
	SetEmail(ctx context.Context, req *v1.SetEmailReq) (res *v1.SetEmailRes, err error)
	TestEmail(ctx context.Context, req *v1.TestEmailReq) (res *v1.TestEmailRes, err error)
	GetMessage(ctx context.Context, req *v1.GetMessageReq) (res *v1.GetMessageRes, err error)
	SetMessage(ctx context.Context, req *v1.SetMessageReq) (res *v1.SetMessageRes, err error)
	TestMessage(ctx context.Context, req *v1.TestMessageReq) (res *v1.TestMessageRes, err error)
	GetAuthConf(ctx context.Context, req *v1.GetAuthConfReq) (res *v1.GetAuthConfRes, err error)
	SetAuthConf(ctx context.Context, req *v1.SetAuthConfReq) (res *v1.SetAuthConfRes, err error)
	CheckPort(ctx context.Context, req *v1.CheckPortReq) (res *v1.CheckPortRes, err error)
	Shell(ctx context.Context, req *v1.ShellReq) (res *v1.ShellRes, err error)
	GetSystemStatus(ctx context.Context, req *v1.GetSystemStatusReq) (res *v1.GetSystemStatusRes, err error)
}
