// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model"
)

type (
	IKernel interface {
		Set(ctx context.Context, m *model.Kernel) error
		Get(ctx context.Context) (m *model.Kernel, err error)
	}
)

var (
	localKernel IKernel
)

func Kernel() IKernel {
	if localKernel == nil {
		panic("implement not found for interface IKernel, forgot register?")
	}
	return localKernel
}

func RegisterKernel(i IKernel) {
	localKernel = i
}
