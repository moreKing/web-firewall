package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"time"

	"server/api/system/v1"
)

func (c *ControllerV1) CheckPort(ctx context.Context, req *v1.CheckPortReq) (res *v1.CheckPortRes, err error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", req.Ip, req.Port), 5*time.Second)
	if err != nil {
		g.Log().Error(ctx, err)
		return &v1.CheckPortRes{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(conn)

	return &v1.CheckPortRes{
		Success: true,
	}, nil
}
