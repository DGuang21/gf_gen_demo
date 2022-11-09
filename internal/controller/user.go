package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	v12 "hello_gf/api/generated/http/echo/v1"
	"hello_gf/internal/consts"
	"hello_gf/internal/service"
)

// 用户管理
var User = cUser{}

type cUser struct{}

func (c cUser) Say(ctx context.Context, req *v12.SayReq) (res *v12.SayRes, err error) {
	if err = service.User().Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(consts.UserLoginUrl)
	return
}
