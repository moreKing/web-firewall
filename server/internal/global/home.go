package global

import (
	"context"
	"fmt"
	"server/internal/dao"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// OnlineUsers 用户各个时间段在线数量
type OnlineUsers struct {
	StartTime  int64 `json:"startTime"`
	UpdateTime int64 `json:"updateTime"`
	Online     []int `json:"online"`
}

var yesterday OnlineUsers
var today OnlineUsers

func GetYesterday(ctx context.Context) *[]int {
	now := time.Now()
	yesterdayTime := now.AddDate(0, 0, -1)
	if len(yesterday.Online) == 0 || yesterday.StartTime < time.Date(yesterdayTime.Year(), yesterdayTime.Month(),
		yesterdayTime.Day(), 0, 0, 0, 0, time.Local).Unix() {
		yesterday.Online = []int{}

		for i := 0; i < 24; i += 2 {
			start := time.Date(yesterdayTime.Year(), yesterdayTime.Month(),
				yesterdayTime.Day(), i, 0, 0, 0, time.Local).Unix()
			end := time.Date(yesterdayTime.Year(), yesterdayTime.Month(),
				yesterdayTime.Day(), i+1, 59, 59, 999, time.Local).Unix()
			count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("(%s = ? AND %s > ? AND %s < ? ) OR (%s = ? AND  %s > ? AND  %s < ? ) OR (%s = ? AND %s < ?)",
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LogoutAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt),
				false, start, end, false, start, start, true, end).Count()
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
			yesterday.Online = append(yesterday.Online, count)
		}

		yesterday.StartTime = time.Date(yesterdayTime.Year(), yesterdayTime.Month(),
			yesterdayTime.Day(), 0, 0, 0, 0, time.Local).Unix()

		yesterday.UpdateTime = now.Unix()

		return &yesterday.Online
	}

	return &yesterday.Online
}

func GetToday(ctx context.Context) *[]int {
	now := time.Now()

	startNum := nextEven(now.Hour())

	if startNum <= 2 {
		start := time.Date(now.Year(), now.Month(),
			now.Day(), 0, 0, 0, 0, time.Local).Unix()
		end := time.Date(now.Year(), now.Month(),
			now.Day(), 1, 59, 59, 999, time.Local).Unix()
		count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("(%s = ? AND %s > ? AND %s < ? ) OR (%s = ? AND  %s > ? AND  %s < ? ) OR (%s = ? AND %s < ?)",
			dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().LoginAt,
			dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LogoutAt, dao.LogLogins.Columns().LoginAt,
			dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt),
			false, start, end, false, start, start, true, end).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}

		return &[]int{count}
	}

	if today.StartTime != time.Date(now.Year(), now.Month(),
		now.Day(), 0, 0, 0, 0, time.Local).Unix() {
		today.Online = []int{}
		for i := 0; i < startNum-2; i += 2 {

			start := time.Date(now.Year(), now.Month(),
				now.Day(), i, 0, 0, 0, time.Local).Unix()
			end := time.Date(now.Year(), now.Month(),
				now.Day(), i+1, 59, 59, 999, time.Local).Unix()
			count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("(%s = ? AND %s > ? AND %s < ? ) OR (%s = ? AND  %s > ? AND  %s < ? ) OR (%s = ? AND %s < ?)",
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LogoutAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt),
				false, start, end, false, start, start, true, end).Count()
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
			today.Online = append(today.Online, count)
			//g.Log().Debug(ctx, "online: ", today.Online)
		}

		today.StartTime = time.Date(now.Year(), now.Month(),
			now.Day(), 0, 0, 0, 0, time.Local).Unix()
		today.UpdateTime = now.Unix()
	}

	// 如果数量不够
	if len(today.Online) < int(startNum/2) {
		for i := len(today.Online) * 2; i < startNum-2; i += 2 {
			start := time.Date(now.Year(), now.Month(),
				now.Day(), i, 0, 0, 0, time.Local).Unix()
			end := time.Date(now.Year(), now.Month(),
				now.Day(), i+1, 59, 59, 999, time.Local).Unix()
			count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("(%s = ? AND %s > ? AND %s < ? ) OR (%s = ? AND  %s > ? AND  %s < ? ) OR (%s = ? AND %s < ?)",
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LogoutAt, dao.LogLogins.Columns().LoginAt,
				dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt),
				false, start, end, false, start, start, true, end).Count()
			if err != nil {
				g.Log().Error(ctx, err)
				return nil
			}
			today.Online = append(today.Online, count)
		}
	}

	// 最新的时间加进去返回
	start := time.Date(now.Year(), now.Month(),
		now.Day(), startNum-2, 0, 0, 0, time.Local).Unix()
	end := time.Date(now.Year(), now.Month(),
		now.Day(), now.Hour(), 59, 59, 999, time.Local).Unix()
	count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("(%s = ? AND %s > ? AND %s < ? ) OR (%s = ? AND  %s > ? AND  %s < ? ) OR (%s = ? AND %s < ?)",
		dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().LoginAt,
		dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LogoutAt, dao.LogLogins.Columns().LoginAt,
		dao.LogLogins.Columns().Online, dao.LogLogins.Columns().LoginAt),
		false, start, end, false, start, start, true, end).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil

	}

	tmp := append(today.Online, count)

	return &tmp
}

// 获取一个偶数
func nextEven(num int) int {
	if num%2 == 0 {
		return num + 2
	} else {
		return num + 1
	}
}
