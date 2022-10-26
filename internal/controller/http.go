package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"

	v1 "hello_gf/api/user/v1"
	"hello_gf/internal/service"
)

func NewHTTPServer(*gcfg.Config, *glog.Logger) *ghttp.Server {
	// 这部分是用户可以自己修改的方法
	// =========
	// 封装的Recovery / 链路追踪 / jwt / 校验链 / logger / 服务注册和发现 / 自定义的中间件在这里做注入
	// 相当于是gf 将上方说到的功能封装成中间件，在这里去执行注入
	// =========
	// 真正启动服务
	s := g.Server()

	// auto generate
	userServiceImpl := &service.User{}

	// auto generate
	v1.RegisterUserHTTPServer(s, userServiceImpl)

	return s
}
