package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "hello_gf/api/generated/time"
)

// Time is the service for time.
type Time struct{}

// Now implements the protobuf.TimeServer interface.
func (s *Time) Now(ctx context.Context, r *v1.NowReq) (*v1.NowRes, error) {
	text := fmt.Sprintf(`%s: %s`, gcmd.GetOpt("node", "default"), gtime.Now().String())
	return &v1.NowRes{Time: text}, nil
}
