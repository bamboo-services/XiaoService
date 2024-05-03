package main

import (
	_ "XiaoService/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"XiaoService/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
