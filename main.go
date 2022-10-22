package main

import (
	_ "hello_gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hello_gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
