package systemconfig

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jordan-wright/email"
	"mime"
	"net/smtp"
	"net/textproto"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/utility/gm"
)

func (s *sSystemConfig) GetEmail() *model.Email {
	return &model.Email{
		Enable:   s.email.Enable,
		SMTP:     s.email.SMTP,
		Port:     s.email.Port,
		Email:    s.email.Email,
		Account:  s.email.Account,
		Protocol: s.email.Protocol,
	}
}

func (s *sSystemConfig) SetEmail(email *model.Email) error {
	if !email.Enable {
		empty := &model.Email{
			Enable:   false,
			SMTP:     "",
			Port:     25,
			Email:    "",
			Account:  "",
			Protocol: 1,
			Password: "",
		}
		_, err := dao.SystemConf.Ctx(context.Background()).Where("id", EmailId).Update(do.SystemConf{Config: empty})
		if err == nil {
			s.email = empty
		}
		return err
	}

	if email.Enable && (email.Port < 1 || email.Port > 365535) {
		return gerror.NewCode(gcode.CodeValidationFailed, fmt.Sprintf("%d 端口不在 1-65535范围", email.Port))
	}

	if !s.email.Enable && email.Enable && email.Password == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "密码不能为空")
	}

	if s.email.Enable && email.Enable && email.Password == "" {
		email.Password = s.email.Password
	}
	tmp := email.Password
	tmp2, err := gm.Sm2Encrypt(email.Password)
	email.Password = tmp2
	_, err = dao.SystemConf.Ctx(context.Background()).Where("id", EmailId).Update(do.SystemConf{Config: email})

	if err == nil {
		s.email = email
		s.email.Password = tmp
	}

	return err
}

func (s *sSystemConfig) SendEmailText(addr, title, content string) error {
	if !s.email.Enable {
		return errors.New("未启用邮箱")
	}
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", s.email.Account, s.email.Email)
	e.To = []string{addr}
	e.Subject = title
	e.Text = []byte(content)
	var err error
	switch s.email.Protocol {
	case 1:
		err = e.Send(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP))

	case 2:
		err = e.SendWithTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})

	case 3:
		err = e.SendWithStartTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})
	}

	return err
}

func (s *sSystemConfig) SendEmailStream(email *model.Email) error {
	if !email.Enable {
		empty := &model.Email{
			Enable:   false,
			SMTP:     "",
			Port:     25,
			Email:    "",
			Account:  "",
			Protocol: 1,
			Password: "",
		}
		_, err := dao.SystemConf.Ctx(context.Background()).Where("id", EmailId).Update(do.SystemConf{Config: empty})
		if err == nil {
			s.email = empty
		}
		return err
	}

	if !s.email.Enable && email.Enable && email.Password == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "密码不能为空")
	}

	if s.email.Enable && email.Enable && email.Password == "" {
		email.Password = s.email.Password
	}
	tmp := email.Password
	tmp2, err := gm.Sm2Encrypt(email.Password)
	email.Password = tmp2
	_, err = dao.SystemConf.Ctx(context.Background()).Where("id", EmailId).Update(do.SystemConf{Config: email})

	if err == nil {
		s.email = email
		s.email.Password = tmp
	}

	return err
}

func (s *sSystemConfig) SendEmailHtml(addr, title, content string) error {
	if !s.email.Enable {
		return errors.New("未启用邮箱")
	}
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", s.email.Account, s.email.Email)
	e.To = []string{addr}
	e.Subject = title
	//e.Text = []byte(content)
	e.HTML = []byte(content)
	var err error
	switch s.email.Protocol {
	case 1:
		err = e.Send(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP))

	case 2:
		err = e.SendWithTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})

	case 3:
		err = e.SendWithStartTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})
	}

	return err
}

func (s *sSystemConfig) SendEmailHtmlZip(addr, title, content, fileName string, f []byte) error {
	if !s.email.Enable {
		return errors.New("未启用邮箱")
	}
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", s.email.Account, s.email.Email)
	e.To = []string{addr}
	e.Subject = title
	//e.Text = []byte(content)
	e.HTML = []byte(content)

	e.Attachments = append(e.Attachments, &email.Attachment{
		Filename:    mime.QEncoding.Encode("utf-8", fileName),
		ContentType: "application/zip",
		Header:      textproto.MIMEHeader{},
		Content:     f,
	})

	var err error
	switch s.email.Protocol {
	case 1:
		err = e.Send(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP))

	case 2:
		err = e.SendWithTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})

	case 3:
		err = e.SendWithStartTLS(fmt.Sprintf("%s:%d", s.email.SMTP, s.email.Port), smtp.PlainAuth("", s.email.Email, s.email.Password, s.email.SMTP),
			&tls.Config{
				InsecureSkipVerify: true,
				ServerName:         s.email.SMTP,
			})
	}

	return err
}
