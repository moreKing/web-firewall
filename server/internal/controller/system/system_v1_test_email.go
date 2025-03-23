package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/global"
	"server/internal/service"

	"server/api/system/v1"
)

func (c *ControllerV1) TestEmail(ctx context.Context, req *v1.TestEmailReq) (res *v1.TestEmailRes, err error) {

	//content, err := g.View().ParseContent(ctx, global.SendEmailTestTpl, g.Map{"username": "超级管理员", "account": service.SystemConfig().GetEmail().Account, "data": g.Slice{
	//	g.Map{"name": "111", "ip": "192.168.1.1", "account": "root", "status": "init"},
	//	g.Map{"name": "222", "ip": "192.168.1.1", "account": "administrator", "status": "init"},
	//	g.Map{"name": "333", "ip": "192.168.1.1", "account": "root", "status": "init"},
	//}})

	content, err := g.View().ParseContent(ctx, global.SendEmailTestTpl, g.Map{"username": "超级管理员", "account": service.SystemConfig().GetEmail().Account})
	if err != nil {
		return nil, err
	}
	err = service.SystemConfig().SendEmailHtml(req.To, "测试邮件", content)
	return nil, err
}
