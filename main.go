package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"hello_gf/internal/cmd"

	_ "hello_gf/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
