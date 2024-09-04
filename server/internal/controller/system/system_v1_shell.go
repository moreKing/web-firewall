package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
	"server/utility/myssh/pty"

	"server/api/system/v1"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *ControllerV1) Shell(ctx context.Context, req *v1.ShellReq) (res *v1.ShellRes, err error) {

	webcon, err := upGrader.Upgrade(ghttp.RequestFromCtx(ctx).Response.Writer, ghttp.RequestFromCtx(ctx).Request, nil)
	if err != nil {
		fmt.Println("升级http 为websoket失败：", err)
		return nil, errors.New("websocket建立失败")
	}

	shell, err := pty.NewShell(webcon, ctx)
	if err != nil {
		return nil, err
	}

	shell.Start()
	return nil, nil
}
