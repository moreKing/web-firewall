package cmd

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "firewall-web",
		Usage: "Firewalld",
		Brief: "start http server",
		Func:  register(),
	}
)
