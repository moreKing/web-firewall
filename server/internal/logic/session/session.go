package session

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"time"
)

type sSession struct {
	onlineUsers map[string]*model.ContextUser
}

func init() {
	service.RegisterSession(New())
	//	从数据库中读取所有在线用户
	var onlines []entity.LogLogins
	ctx := gctx.New()
	err := dao.LogLogins.Ctx(ctx).Where(dao.LogLogins.Columns().Online, true).Scan(&onlines)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, len(onlines))

	for _, online := range onlines {
		if online.LogoutAt < time.Now().Unix() {
			// 更新数据库中的数据
			_, err = dao.LogLogins.Ctx(ctx).Data(do.LogLogins{
				Online: false,
				Log:    "超时退出",
			}).Where(dao.LogLogins.Columns().Uuid, online.Uuid).Update()
			if err != nil {
				g.Log().Error(ctx, err)
			}
			continue
		}
		service.Session().AddOnlineUser(&model.ContextUser{
			ID:             online.UserId,
			Authentication: online.Uuid,
			LoginName:      online.Loginname,
			Password:       "",
			ClientIP:       online.ClientIp,
			CreatedAt:      online.LoginAt,
			ExpiredAt:      online.LogoutAt,
		})
	}
}

func New() service.ISession {
	return &sSession{
		onlineUsers: make(map[string]*model.ContextUser),
	}
}

// AddOnlineUser 登录成功后调用这个方法进行token注册
func (s *sSession) AddOnlineUser(m *model.ContextUser) bool {
	s.onlineUsers[m.Authentication] = m
	return false
}

// RemoveOnlineUser 退出登录调用此方法取消登录状态
func (s *sSession) RemoveOnlineUser(ctx context.Context, authentication, description string, logoutAt int64) bool {
	_, err := dao.LogLogins.Ctx(ctx).Data(do.LogLogins{
		Online:   false,
		Log:      description,
		LogoutAt: logoutAt,
	}).Where(dao.LogLogins.Columns().Uuid, authentication).Update()
	if err != nil {
		g.Log().Error(ctx, err)
	}
	delete(s.onlineUsers, authentication)
	return true
}

// GetOnlineUser 校验token是否存在，存在返回对象
func (s *sSession) GetOnlineUser(authentication string) (*model.ContextUser, bool) {
	user, ok := s.onlineUsers[authentication]

	return user, ok
}

func (s *sSession) GetUser(ctx context.Context) *entity.Users {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() {
			var user *entity.Users
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

// RemoveUser removes user rom session.
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}

// ClearExpiredSession 清理过期会话
func (s *sSession) ClearExpiredSession(ctx context.Context) {
	for _, user := range s.onlineUsers {
		if user.ExpiredAt < time.Now().Unix() {
			g.Log().Debug(ctx, "自动清理过期会话 ", user)
			// 更新数据库中的数据
			s.RemoveOnlineUser(ctx, user.Authentication, "超时退出", int64(user.ExpiredAt))
		} else {
			// 没有过期的 更新退出时间，为了服务重启后 读取在线用户服务
			_, err := dao.LogLogins.Ctx(ctx).Data(do.LogLogins{
				LogoutAt: user.ExpiredAt,
			}).Where(dao.LogLogins.Columns().Uuid, user.Authentication).Update()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}
	}
}
