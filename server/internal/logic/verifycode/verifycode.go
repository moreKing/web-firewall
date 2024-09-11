package verifycode

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xlzd/gotp"
	"math/rand/v2"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/service"
	"sync"
	"time"
)

/**
双因素认证 验证码处理
*/

type sCodeServer struct {
	m     sync.Mutex
	codes map[string]*model.VerificationCode
}

func init() {
	service.RegisterCodeServer(New())
}

func New() service.ICodeServer {
	return &sCodeServer{
		m:     sync.Mutex{},
		codes: make(map[string]*model.VerificationCode),
	}
}

func (s *sCodeServer) VerifyCode(ctx context.Context, id int64, token, code, ip string) bool {
	st, ok := s.codes[token]
	if !ok {
		return false
	}

	defer delete(s.codes, token)

	if st.Code == code && st.ExpiredAt > time.Now().Unix() && st.ClientIP == ip && st.UserId == id {
		return true
	}

	return false
}

func (s *sCodeServer) CreateCode(ctx context.Context, id int64, ip string, offset int) (*model.VerificationCode, error) {
	s.m.Lock()
	defer s.m.Unlock()
	// 同一个用户 1小时内不能创建相同的验证码，创建验证码后校验验证码是否被使用
	var codeStr string
	for true {
		codeStr = s.generateVerificationCode(6)
		for _, code := range s.codes {
			if code.Code == codeStr && code.UserId == id {
				continue
			}
		}
		// 去查数据库近一个小时同一个用户有没有登陆成功过的
		count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("%s > ? AND  %s = ? AND %s = ? AND %s  = ?", dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().UserId, dao.LogLogins.Columns().Success, dao.LogLogins.Columns().TotpCode), time.Now().Add(time.Hour*-1).Unix(), id, true, codeStr).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, err
		}
		if count > 0 {
			continue
		}
		break
	}

	// 验证码没有使用过 生成uuid，并添加到数组中
	vcode := &model.VerificationCode{
		Token:     gonanoid.Must(),
		ClientIP:  ip,
		UserId:    id,
		Code:      codeStr,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(offset)).Unix(),
	}

	s.codes[vcode.Token] = vcode

	return vcode, nil
}

func (s *sCodeServer) generateVerificationCode(length int) string {
	// 数字跟大写字母 不包含小写 没有字母 O I
	digits := "0123456789"
	uppercases := "ABCDEFGHJKLMNPQRSTUVWXYZ" // 不要字母O I 有歧义
	//lowerCases := "abcdefghijklmnopqrstuvwxyz"
	//lowerCases := ""+ lowerCases
	all := uppercases + digits
	buf := make([]byte, length)
	// 随机生成
	for i := 0; i < length; i++ {
		buf[i] = all[rand.IntN(len(all))]
	}

	// 随机排列
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func (s *sCodeServer) VerifyTotp(ctx context.Context, code, secret string, userid int64) error {
	s.m.Lock()
	defer s.m.Unlock()
	//最近一小时有没有登陆过
	count, err := dao.LogLogins.Ctx(ctx).Where(fmt.Sprintf("%s > ? AND  %s = ? AND %s = ? AND %s = ?", dao.LogLogins.Columns().LoginAt, dao.LogLogins.Columns().UserId, dao.LogLogins.Columns().Success, dao.LogLogins.Columns().TotpCode), time.Now().Add(time.Hour*-1).Unix(), userid, true, code).Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count > 0 {
		return errors.New("此令牌1小时内使用过")
	}

	offset := service.SystemConfig().GetAuthConf().TotpOffset
	t := time.Now().Unix()
	for i := offset * -1; i <= offset; i++ {
		if ok := gotp.NewDefaultTOTP(secret).Verify(code, t+int64(i*30)); ok {
			return nil
		}
	}
	return errors.New("令牌有效期内验证失败")
}

func (s *sCodeServer) RemoveExpireCode() {

	nowTime := time.Now().Unix()
	for _, code := range s.codes {
		if code.ExpiredAt < nowTime {
			g.Log().Debug(context.Background(), "RemoveExpireCode ", code)
			delete(s.codes, code.Token)
		}
	}
}

// 复用一下这个作为websocket的token，懒得再写一个了
func (s *sCodeServer) CreateWebsocketToken(ctx context.Context, id int64, ip string, offset int) (*model.VerificationCode, error) {
	s.m.Lock()
	defer s.m.Unlock()
	// 同一个用户 1小时内不能创建相同的验证码，创建验证码后校验验证码是否被使用

	// 验证码没有使用过 生成uuid，并添加到数组中
	vcode := &model.VerificationCode{
		Token:     gonanoid.Must(),
		ClientIP:  ip,
		UserId:    id,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: time.Now().Add(time.Minute * time.Duration(offset)).Unix(),
	}

	s.codes[vcode.Token] = vcode

	return vcode, nil
}

func (s *sCodeServer) VerifyWebsocketToken(ctx context.Context, secret, clientIP string) (userId int64, err error) {
	s.m.Lock()
	defer s.m.Unlock()

	code, ok := s.codes[secret]
	if !ok {
		return 0, errors.New("无效的请求")
	}

	if code.ClientIP != clientIP {
		return 0, errors.New("token与地址不符")
	}
	delete(s.codes, secret)

	return code.UserId, nil

}
