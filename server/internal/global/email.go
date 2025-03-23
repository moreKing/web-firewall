package global

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"path"
)

var SendEmailTestTpl string

var SendEmailCodeTpl string

func init() {
	getwd, err := os.Getwd()
	if err != nil {
		g.Log().Fatal(context.TODO(), "getwd err", err)
		return
	}
	// 测试邮件模板
	file, err := os.ReadFile(path.Join(getwd, "/resource/template/emailTestTpl.html"))
	if err != nil {
		g.Log().Fatal(context.Background(), "读emailTestTpl文件失败", err)
		return
	}
	SendEmailTestTpl = string(file)

	file2, err := os.ReadFile(path.Join(getwd, "/resource/template/sendEmailCodeTpl.html"))
	if err != nil {
		g.Log().Fatal(context.Background(), "读sendEmailCodeTpl文件失败", err)
		return
	}
	SendEmailCodeTpl = string(file2)

}
