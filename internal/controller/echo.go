package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	v1 "hello_gf/api/generated/echo"
)

// Echo is the service for echo.
type Echo struct{}

// Say implements the protobuf.EchoServer interface.
func (s *Echo) Say(ctx context.Context, r *v1.SayReq) (*v1.SayRes, error) {
	g.Log().Print(ctx, "Received:", r.Content)
	text := fmt.Sprintf(`%s: > %s`, gcmd.GetOpt("node", "default"), r.Content)
	return &v1.SayRes{Content: text}, nil
}
