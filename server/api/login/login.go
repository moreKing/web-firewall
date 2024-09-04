// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"server/api/login/v1"
)

type ILoginV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
}
