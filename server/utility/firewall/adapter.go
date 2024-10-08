package firewall

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

type Firewall interface {
	Flush(ctx context.Context) error
}

func New() Firewall {
	// 读取配置文件，根据写入的配置来创建对应防火墙
	t := g.Cfg().MustGet(context.Background(), "firewall.type", "nftables").String()

	if strings.TrimSpace(t) == "iptables" {
		return &Iptables{}
	}

	return NewNftables()
}
