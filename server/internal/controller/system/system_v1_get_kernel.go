package system

import (
	"context"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) GetKernel(ctx context.Context, req *v1.GetKernelReq) (res *v1.GetKernelRes, err error) {
	conf, err := service.Kernel().Get(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetKernelRes{
		Kernel: conf,
	}, nil
}
