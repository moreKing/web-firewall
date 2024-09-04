package main

import (
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	_ "server/internal/packed"

	_ "server/internal/logic"


	"github.com/gogf/gf/v2/os/gctx"

	"server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
