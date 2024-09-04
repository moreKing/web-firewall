// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package public

import (
	"context"

	"server/api/public/v1"
)

type IPublicV1 interface {
	GetPasswordComplex(ctx context.Context, req *v1.GetPasswordComplexReq) (res *v1.GetPasswordComplexRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error)
	SetUserPassword(ctx context.Context, req *v1.SetUserPasswordReq) (res *v1.SetUserPasswordRes, err error)
	SetProfile(ctx context.Context, req *v1.SetProfileReq) (res *v1.SetProfileRes, err error)
}
