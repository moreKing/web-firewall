// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package policy

import (
	"context"

	"server/api/policy/v1"
)

type IPolicyV1 interface {
	GetInputPolicy(ctx context.Context, req *v1.GetInputPolicyReq) (res *v1.GetInputPolicyRes, err error)
	AddInputPolicy(ctx context.Context, req *v1.AddInputPolicyReq) (res *v1.AddInputPolicyRes, err error)
	ReplaceInputPolicy(ctx context.Context, req *v1.ReplaceInputPolicyReq) (res *v1.ReplaceInputPolicyRes, err error)
	ChangeInputPolicyPosition(ctx context.Context, req *v1.ChangeInputPolicyPositionReq) (res *v1.ChangeInputPolicyPositionRes, err error)
	DeleteInputPolicy(ctx context.Context, req *v1.DeleteInputPolicyReq) (res *v1.DeleteInputPolicyRes, err error)
	GetInputLimitPolicy(ctx context.Context, req *v1.GetInputLimitPolicyReq) (res *v1.GetInputLimitPolicyRes, err error)
	AddInputLimitPolicy(ctx context.Context, req *v1.AddInputLimitPolicyReq) (res *v1.AddInputLimitPolicyRes, err error)
	ReplaceInputLimitPolicy(ctx context.Context, req *v1.ReplaceInputLimitPolicyReq) (res *v1.ReplaceInputLimitPolicyRes, err error)
	ChangeInputLimitPolicyPosition(ctx context.Context, req *v1.ChangeInputLimitPolicyPositionReq) (res *v1.ChangeInputLimitPolicyPositionRes, err error)
	DeleteInputLimitPolicy(ctx context.Context, req *v1.DeleteInputLimitPolicyReq) (res *v1.DeleteInputLimitPolicyRes, err error)
	GetOutputPolicy(ctx context.Context, req *v1.GetOutputPolicyReq) (res *v1.GetOutputPolicyRes, err error)
	AddOutputPolicy(ctx context.Context, req *v1.AddOutputPolicyReq) (res *v1.AddOutputPolicyRes, err error)
	ReplaceOutputPolicy(ctx context.Context, req *v1.ReplaceOutputPolicyReq) (res *v1.ReplaceOutputPolicyRes, err error)
	ChangeOutputPolicyPosition(ctx context.Context, req *v1.ChangeOutputPolicyPositionReq) (res *v1.ChangeOutputPolicyPositionRes, err error)
	DeleteOutputPolicy(ctx context.Context, req *v1.DeleteOutputPolicyReq) (res *v1.DeleteOutputPolicyRes, err error)
	GetOutputLimitPolicy(ctx context.Context, req *v1.GetOutputLimitPolicyReq) (res *v1.GetOutputLimitPolicyRes, err error)
	AddOutputLimitPolicy(ctx context.Context, req *v1.AddOutputLimitPolicyReq) (res *v1.AddOutputLimitPolicyRes, err error)
	ReplaceOutputLimitPolicy(ctx context.Context, req *v1.ReplaceOutputLimitPolicyReq) (res *v1.ReplaceOutputLimitPolicyRes, err error)
	ChangeOutputLimitPolicyPosition(ctx context.Context, req *v1.ChangeOutputLimitPolicyPositionReq) (res *v1.ChangeOutputLimitPolicyPositionRes, err error)
	DeleteOutputLimitPolicy(ctx context.Context, req *v1.DeleteOutputLimitPolicyReq) (res *v1.DeleteOutputLimitPolicyRes, err error)
}
