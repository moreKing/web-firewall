package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/api/hello/v1"
	"server/utility/nftables"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	nftables.CreateTable(&nftables.Table{
		Name:    "hello",
		Comment: "测试",
		Family:  "ip",
	})

	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
