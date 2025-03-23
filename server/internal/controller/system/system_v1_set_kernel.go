package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) SetKernel(ctx context.Context, req *v1.SetKernelReq) (res *v1.SetKernelRes, err error) {
	err = service.Kernel().Set(ctx, req.Kernel)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
