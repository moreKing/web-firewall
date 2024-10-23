package firewall

import (
	"context"
	"os/exec"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

type Firewall interface {
	Flush(ctx context.Context) error
}

func New() Firewall {
	// 读取配置文件，根据写入的配置来创建对应防火墙
	t := g.Cfg().MustGet(context.Background(), "firewall.type", "auto").String()

	if strings.TrimSpace(t) == "iptables" {
		return &Iptables{}
	}

	if strings.TrimSpace(t) == "nftables" {
		return NewNftables()
	}

	// 判断防火墙类型
	tp := g.Cfg().MustGet(context.Background(), "firewall.typePriority", "nftables").String()

	if strings.TrimSpace(tp) == "iptables" {
		// 优先判断是否是 iptables
		if isFirewalld("iptables", "-V") {
			return &Iptables{}
		}
	}

	if isFirewalld("nft", "-v") {
		return NewNftables()
	}

	if isFirewalld("iptables", "-V") {
		return &Iptables{}
	}

	panic("无适配的防火墙类型")
}

func isFirewalld(t string, arg string) bool {
	output, err := exec.Command(t, arg).CombinedOutput()
	if err != nil {
		g.Log().Error(context.Background(), "防火墙类型检测失败", t, string(output))
		return false
	}

	g.Log().Info(context.Background(), "防火墙类型检测成功：", string(output))
	return true
}
