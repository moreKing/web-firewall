// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
)

type (
	ITimer interface {
		// Add 立刻执行, interval单位：天
		Add(ctx context.Context, tag string, interval int, f func()) error
		// AddStartAt 首次执行时间，与Add 不同的是 Add是立刻执行，本函数可以指定首次执行的时间，时间必须是未来的时间
		AddStartAt(ctx context.Context, tag string, interval int, timeAt time.Time, f func()) error
		// AddStartTimestampAt timeAt:首次执行时间毫秒时间戳，时间必须是未来的时间
		AddStartTimestampAt(ctx context.Context, tag string, interval int, timeAt int64, f func()) error
		Remove(tag string)
	}
)

var (
	localTimer ITimer
)

func Timer() ITimer {
	if localTimer == nil {
		panic("implement not found for interface ITimer, forgot register?")
	}
	return localTimer
}

func RegisterTimer(i ITimer) {
	localTimer = i
}
