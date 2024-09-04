// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package home

import (
	"context"

	"server/api/home/v1"
)

type IHomeV1 interface {
	GetHome(ctx context.Context, req *v1.GetHomeReq) (res *v1.GetHomeRes, err error)
}
