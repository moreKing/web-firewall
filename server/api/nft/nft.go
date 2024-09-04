// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package nft

import (
	"context"

	"server/api/nft/v1"
)

type INftV1 interface {
	GetPolicy(ctx context.Context, req *v1.GetPolicyReq) (res *v1.GetPolicyRes, err error)
	AddPolicy(ctx context.Context, req *v1.AddPolicyReq) (res *v1.AddPolicyRes, err error)
	ReplacePolicy(ctx context.Context, req *v1.ReplacePolicyReq) (res *v1.ReplacePolicyRes, err error)
	ChangePolicyPosition(ctx context.Context, req *v1.ChangePolicyPositionReq) (res *v1.ChangePolicyPositionRes, err error)
	DeletePolicy(ctx context.Context, req *v1.DeletePolicyReq) (res *v1.DeletePolicyRes, err error)
}
