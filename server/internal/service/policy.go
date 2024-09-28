// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPolicy interface {
		Flush(ctx context.Context) error
	}
)

var (
	localPolicy IPolicy
)

func Policy() IPolicy {
	if localPolicy == nil {
		panic("implement not found for interface IPolicy, forgot register?")
	}
	return localPolicy
}

func RegisterPolicy(i IPolicy) {
	localPolicy = i
}
