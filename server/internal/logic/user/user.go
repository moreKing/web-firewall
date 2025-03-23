package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"server/internal/dao"
	"server/internal/model/do"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility/gm"
	"server/utility/password"
	"time"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// SetNativePassword 设置用户本地登录密码
func (s *sUser) SetNativePassword(ctx context.Context, userId int64, newPassword string) (err error) {
	// 0. 校验密码是否符合规则
	// 0.1 获取全局密码强度
	pc := service.SystemConfig().GetUserPasswordComplex()
	if pc.Length > len(newPassword) {
		return gerror.NewCode(gcode.CodeValidationFailed, fmt.Sprintf("新密码长度小于%d", pc.Length))
	}

	comp := password.PasswordComplexity(newPassword)

	switch pc.Complex {
	case 2:
		if !(comp.Digit > 0 && (comp.Upper > 0 || comp.Lowercase > 0)) {
			return gerror.NewCode(gcode.CodeValidationFailed, "新密码至少包含数字、字母")
		}
	case 3:
		if !(comp.Digit > 0 && (comp.Upper > 0 || comp.Lowercase > 0) && comp.Other > 0) {
			return gerror.NewCode(gcode.CodeValidationFailed, "新密码至少包含字母、数字、特殊字符 ")
		}
	case 4:
		if !(comp.Digit > 0 && comp.Upper > 0 && comp.Lowercase > 0 && comp.Other > 0) {
			return gerror.NewCode(gcode.CodeValidationFailed, "新密码至少包含大写字母、小写字母、数字、特殊字符")
		}
	}

	// 1. 随机生成slat
	slat, _ := gonanoid.New()
	// 2. 生成sm3加密密码
	pass := gm.SM3(newPassword, slat)
	g.Log().Debug(ctx, fmt.Sprintf("slat:%s pass sm3:%s", slat, pass))
	// 3. 保存
	_, err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, userId).Fields(dao.Users.Columns().Slat, dao.Users.Columns().Password, dao.Users.Columns().PwdUpdateAt, dao.Users.Columns().UpdatedAt).Update(&entity.Users{Id: userId, Slat: slat, Password: pass, PwdUpdateAt: time.Now().Unix()})
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID 查询指定用户信息
func (s *sUser) GetUserByID(ctx context.Context, id int64) (user *entity.Users, err error) {
	//var user *entity.Users
	err = dao.Users.Ctx(ctx).Where(do.Users{
		Id: id,
	}).Scan(&user)
	if err != nil || user == nil {
		err = gerror.Newf("不存在的用户id : %s", id)
		return
	}
	return
}
