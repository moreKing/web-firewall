package timer

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"server/internal/service"
	"time"
)

//var Cron gocron.Scheduler

// 定时任务管理器
type sTimer struct {
	cron gocron.Scheduler
}

func init() {
	service.RegisterTimer(New())
	g.Log().Debug(context.Background(), "创建注册定时器完成")
}

func New() service.ITimer {
	s, err := gocron.NewScheduler()
	timer := &sTimer{
		cron: s,
	}
	if err != nil {
		// handle error
		g.Log().Fatal(context.Background(), "无法创建定时器 ", err)
	}

	// 10min 清理过期登录
	_, err = timer.cron.NewJob(
		gocron.CronJob("*/10 * * * *", false),
		gocron.NewTask(
			func() {
				// do things
				service.Session().ClearExpiredSession(gctx.New())
			},
		),
		gocron.WithTags("ClearExpiredSession"),
	)

	// 30min 清理过期验证码
	_, err = timer.cron.NewJob(
		gocron.CronJob("*/30 * * * *", false),
		gocron.NewTask(
			func() {
				// do things
				service.CodeServer().RemoveExpireCode()
			},
		),
		gocron.WithTags("RemoveExpireCode"),
	)

	//开始运行定时器
	timer.cron.Start()

	return timer
}

// Add 立刻执行, interval单位：天
func (s *sTimer) Add(ctx context.Context, tag string, interval int, f func()) error {
	_, err := s.cron.NewJob(
		gocron.DurationJob(time.Duration(interval)*time.Hour),
		gocron.NewTask(f),
		gocron.WithTags(tag),
	)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// AddStartAt 首次执行时间，与Add 不同的是 Add是立刻执行，本函数可以指定首次执行的时间，时间必须是未来的时间
func (s *sTimer) AddStartAt(ctx context.Context, tag string, interval int, timeAt time.Time, f func()) error {
	_, err := s.cron.NewJob(
		gocron.DurationJob(time.Duration(interval)*time.Hour*24),
		gocron.NewTask(f),
		gocron.WithTags(tag),
		gocron.WithStartAt(gocron.WithStartDateTime(timeAt)),
	)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// AddStartTimestampAt timeAt:首次执行时间毫秒时间戳，时间必须是未来的时间
func (s *sTimer) AddStartTimestampAt(ctx context.Context, tag string, interval int, timeAt int64, f func()) error {
	now := time.Now().UnixMilli()
	if timeAt < now {
		for {
			timeAt = timeAt + int64(interval*24*3600*1000)
			if timeAt > now {
				break
			}
		}
	}

	startTime := time.UnixMilli(timeAt)
	_, err := s.cron.NewJob(
		gocron.DurationJob(time.Duration(interval)*time.Hour*24),
		gocron.NewTask(f),
		gocron.WithTags(tag),
		gocron.WithStartAt(gocron.WithStartDateTime(startTime)),
	)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (s *sTimer) Remove(tag string) {
	s.cron.RemoveByTags(tag)
}
