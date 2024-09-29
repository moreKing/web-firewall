// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package route

import (
	"context"

	"server/api/route/v1"
)

type IRouteV1 interface {
	GetDnatPolicy(ctx context.Context, req *v1.GetDnatPolicyReq) (res *v1.GetDnatPolicyRes, err error)
	AddDnatPolicy(ctx context.Context, req *v1.AddDnatPolicyReq) (res *v1.AddDnatPolicyRes, err error)
	ReplaceDnatPolicy(ctx context.Context, req *v1.ReplaceDnatPolicyReq) (res *v1.ReplaceDnatPolicyRes, err error)
	ChangeDnatPolicyPosition(ctx context.Context, req *v1.ChangeDnatPolicyPositionReq) (res *v1.ChangeDnatPolicyPositionRes, err error)
	DeleteDnatPolicy(ctx context.Context, req *v1.DeleteDnatPolicyReq) (res *v1.DeleteDnatPolicyRes, err error)
	GetForwardPolicy(ctx context.Context, req *v1.GetForwardPolicyReq) (res *v1.GetForwardPolicyRes, err error)
	AddForwardPolicy(ctx context.Context, req *v1.AddForwardPolicyReq) (res *v1.AddForwardPolicyRes, err error)
	ReplaceForwardPolicy(ctx context.Context, req *v1.ReplaceForwardPolicyReq) (res *v1.ReplaceForwardPolicyRes, err error)
	ChangeForwardPolicyPosition(ctx context.Context, req *v1.ChangeForwardPolicyPositionReq) (res *v1.ChangeForwardPolicyPositionRes, err error)
	DeleteForwardPolicy(ctx context.Context, req *v1.DeleteForwardPolicyReq) (res *v1.DeleteForwardPolicyRes, err error)
	GetLimitPolicy(ctx context.Context, req *v1.GetLimitPolicyReq) (res *v1.GetLimitPolicyRes, err error)
	AddLimitPolicy(ctx context.Context, req *v1.AddLimitPolicyReq) (res *v1.AddLimitPolicyRes, err error)
	ReplaceLimitPolicy(ctx context.Context, req *v1.ReplaceLimitPolicyReq) (res *v1.ReplaceLimitPolicyRes, err error)
	ChangeLimitPolicyPosition(ctx context.Context, req *v1.ChangeLimitPolicyPositionReq) (res *v1.ChangeLimitPolicyPositionRes, err error)
	DeleteLimitPolicy(ctx context.Context, req *v1.DeleteLimitPolicyReq) (res *v1.DeleteLimitPolicyRes, err error)
	GetSnatPolicy(ctx context.Context, req *v1.GetSnatPolicyReq) (res *v1.GetSnatPolicyRes, err error)
	AddSnatPolicy(ctx context.Context, req *v1.AddSnatPolicyReq) (res *v1.AddSnatPolicyRes, err error)
	ReplaceSnatPolicy(ctx context.Context, req *v1.ReplaceSnatPolicyReq) (res *v1.ReplaceSnatPolicyRes, err error)
	ChangeSnatPolicyPosition(ctx context.Context, req *v1.ChangeSnatPolicyPositionReq) (res *v1.ChangeSnatPolicyPositionRes, err error)
	DeleteSnatPolicy(ctx context.Context, req *v1.DeleteSnatPolicyReq) (res *v1.DeleteSnatPolicyRes, err error)
}
