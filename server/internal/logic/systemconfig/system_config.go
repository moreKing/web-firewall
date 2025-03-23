package systemconfig

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/service"
	"server/utility/gm"
	"slices"
	"strings"
	"time"
)

const (
	PasswordComplexId    = 1 // 用户登录密码复杂度
	AuthenticateId       = 2 // 用户登录认证相关配置
	WebTimeoutId         = 3 // web超时时间
	MessageId            = 4 // 短信认证
	EmailId              = 5 // 邮箱配置
	FileServer1          = 6 // 文件服务器1
	FileServer2          = 7 // 文件服务器2
	AccountExceptionRule = 8 // 异常账号规则

)

var FileServer = map[int64]int{
	1: FileServer1,
	2: FileServer2,
}

// 系统配置
type sSystemConfig struct {
	passwordComplex      *model.PasswordComplex
	authenticate         *model.AuthenticateConf
	webSession           *model.WebSession
	message              *model.Message
	email                *model.Email
	fileServer           map[int64]*model.FileServer
	accountExceptionRule *model.AccountExceptionRule
}

func init() {
	service.RegisterSystemConfig(New())
}

func New() service.ISystemConfig {
	pc := &model.SystemConfPasswordComplex{}
	err := dao.SystemConf.Ctx(context.Background()).Where("id", PasswordComplexId).Scan(pc)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}

	g.Log().Debug(context.Background(), "New  PasswordComplex success ", pc)
	auth := &model.SystemConfAuthenticate{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", AuthenticateId).Scan(auth)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}
	webSession := &model.SystemConfWebSession{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", WebTimeoutId).Scan(webSession)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}
	message := &model.SystemConfMessage{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", MessageId).Scan(message)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}
	tempEmail := &model.SystemConfEmail{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", EmailId).Scan(tempEmail)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}

	// 密码解密
	if tempEmail.Config.Enable && tempEmail.Config.Password != "" {
		tempEmail.Config.Password, err = gm.Sm2Decode(tempEmail.Config.Password)
		if err != nil {
			g.Log().Fatal(context.Background(), "New  System Email Config err:", err)
		}
	}

	fileServer1 := &model.SystemConfFileServer{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", FileServer1).Scan(fileServer1)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}
	fileServer1.Config.Password, _ = gm.Sm2Decode(fileServer1.Config.Password)
	fileServer1.Config.SecretKey, _ = gm.Sm2Decode(fileServer1.Config.SecretKey)

	fileServer2 := &model.SystemConfFileServer{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", FileServer2).Scan(fileServer2)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}

	fileServer2.Config.Password, _ = gm.Sm2Decode(fileServer2.Config.Password)
	fileServer2.Config.SecretKey, _ = gm.Sm2Decode(fileServer2.Config.SecretKey)

	accountExceptionRule := &model.SystemConfAccountExceptionRule{}
	err = dao.SystemConf.Ctx(context.Background()).Where("id", AccountExceptionRule).Scan(accountExceptionRule)
	if err != nil {
		g.Log().Fatal(context.Background(), "New  SystemConfig err:", err)
	}

	return &sSystemConfig{
		passwordComplex:      &pc.Config,
		authenticate:         &auth.Config,
		webSession:           &webSession.Config,
		message:              &message.Config,
		email:                &tempEmail.Config,
		fileServer:           map[int64]*model.FileServer{1: &fileServer1.Config, 2: &fileServer2.Config},
		accountExceptionRule: &accountExceptionRule.Config,
	}
}

func (s *sSystemConfig) GetUserPasswordComplex() *model.PasswordComplex {
	return s.passwordComplex
}

func (s *sSystemConfig) SetUserPasswordComplex(ctx context.Context, conf *model.PasswordComplex) error {
	//数据校验
	if err := g.Validator().Data(conf).Run(ctx); err != nil {
		return err
	}
	// 更新数据
	_, err := dao.SystemConf.Ctx(ctx).Fields(dao.SystemConf.Columns().Config).Where("id", PasswordComplexId).Update(&model.SystemConfPasswordComplex{
		Config: *conf,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	s.passwordComplex = conf
	return nil
}

func (s *sSystemConfig) GetWebSession() *model.WebSession {

	return s.webSession
}

func (s *sSystemConfig) SetWebSession(ws *model.WebSession) error {
	if ws.Timeout > 1440 || ws.Timeout < 10 {
		return errors.New("web超时必须在10-1440之间")
	}
	_, err := dao.SystemConf.Ctx(context.Background()).Where("id", WebTimeoutId).Update(do.SystemConf{Config: ws})
	if err == nil {
		s.webSession = ws
	}
	return err
}

func (s *sSystemConfig) GetMessage() *model.Message {
	return s.message
}

func (s *sSystemConfig) SetMessage(messageConf *model.Message) error {

	if messageConf.State == 1 {
		if !slices.Contains(consts.MessageRequestMethod, messageConf.Method) {
			return errors.New("请求方法必须是POST,GET,PUT,PATCH 中的一种")
		}
		if !slices.Contains(consts.MessageRequestParameterType, messageConf.EncType) {
			return errors.New("请求方法必须是json,form-data中的一种")
		}
	}

	_, err := dao.SystemConf.Ctx(context.Background()).Where("id", MessageId).Update(do.SystemConf{Config: messageConf})
	if err == nil {
		s.message = messageConf
	}
	return err
}

func (s *sSystemConfig) SendMessage(ctx context.Context, code, to string) error {
	err := s.sendMessage(ctx, code, to)
	if err != nil {
		g.Log().Error(context.Background(), "短信发送失败，错误：", err)
		return err
	}
	return nil
}

func (s *sSystemConfig) sendMessage(ctx context.Context, code, to string) error {
	if s.message.State == 0 {
		return errors.New("短信未启用")
	}
	// 处理变量
	content := strings.ReplaceAll(s.message.Content, "{code}", code)
	content = strings.ReplaceAll(content, "{validity}", fmt.Sprintf("%d", s.authenticate.MessageOffset))

	// https 证书错误处理
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 不校验 证书
	}
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}

	var res *http.Response

	// 判断是否是GET
	if strings.ToUpper(s.message.Method) == "GET" {
		values := url.Values{}

		for _, parameter := range s.message.Parameters {
			if parameter.Value == "{mobile}" {
				values.Add(parameter.Key, to)
				continue
			}
			if parameter.Value == "{content}" {
				values.Add(parameter.Key, content)
				continue
			}

			values.Add(parameter.Key, fmt.Sprintf("%v", parameter.Value))
		}
		//values.Add("username", "你好啊")
		req, _ := http.NewRequest("GET", fmt.Sprintf("%s?%s", s.message.URL, values.Encode()), nil)
		res, _ = client.Do(req)

	} else {
		// 判断是否是 form
		if s.message.EncType != "json" {
			body := new(bytes.Buffer)
			w := multipart.NewWriter(body)

			for _, parameter := range s.message.Parameters {
				if parameter.Value == "{mobile}" {
					_ = w.WriteField(parameter.Key, to)
					continue
				}
				if parameter.Value == "{content}" {
					_ = w.WriteField(parameter.Key, content)
					continue
				}
				_ = w.WriteField(parameter.Key, fmt.Sprintf("%v", parameter.Value))
			}
			_ = w.Close()
			req, _ := http.NewRequest(strings.ToUpper(s.message.Method), s.message.URL, body)
			fmt.Println(w.FormDataContentType())
			req.Header.Set("Content-Type", w.FormDataContentType())
			res, _ = client.Do(req)
		} else {
			reqBody := make(map[string]any)
			// json数据
			for _, parameter := range s.message.Parameters {
				if parameter.Value == "{mobile}" {
					reqBody[parameter.Key] = to
					continue
				}
				if parameter.Value == "{content}" {
					reqBody[parameter.Key] = content
					continue
				}
				reqBody[parameter.Key] = parameter.Value
			}

			userBytes, err := json.Marshal(reqBody)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
			payload := bytes.NewReader(userBytes)
			req, err := http.NewRequest(s.message.Method, s.message.URL, payload)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
			req.Header.Add("Content-Type", "application/json")
			res, err = client.Do(req)
			if err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		}

	}

	// 统一处理返回结果
	defer func(Body io.ReadCloser) {
		if Body != nil {
			err := Body.Close()
			if err != nil {
				g.Log().Error(ctx, err)

			}
		}
	}(res.Body)

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(res.Body)
	return errors.New(string(body))

}

func (s *sSystemConfig) GetAuthConf() *model.AuthenticateConf {
	return s.authenticate
}

func (s *sSystemConfig) SetAuthConf(ctx context.Context, a *model.AuthenticateConf) error {
	_, err := dao.SystemConf.Ctx(ctx).Where("id", AuthenticateId).Update(do.SystemConf{Config: a})
	if err == nil {
		s.authenticate = a
	}
	return err
}

func (s *sSystemConfig) GetAccountExceptionRule() *model.AccountExceptionRule {

	return s.accountExceptionRule
}

func (s *sSystemConfig) SetAccountExceptionRule(ctx context.Context, rule *model.AccountExceptionRule) error {

	if rule.PwdWeak.Length < rule.PwdWeak.Digit+rule.PwdWeak.Lower+rule.PwdWeak.Upper+rule.PwdWeak.Other {
		return errors.New("密码长度不能小于各类型之和")
	}

	_, err := dao.SystemConf.Ctx(ctx).Where(dao.SystemConf.Columns().Id, AccountExceptionRule).Update(do.SystemConf{Config: rule})
	if err == nil {
		s.accountExceptionRule = rule
	}

	return nil
}
