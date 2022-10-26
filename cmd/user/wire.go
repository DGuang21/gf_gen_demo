//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package user

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/google/wire"

	"hello_gf/internal/controller"
)

// wireApp init kratos application.
func wireApp(*gcfg.Config, glog.Logger) (*ghttp.Server, func(), error) {
	// 这里定义的是需要依赖wire生成的ProviderSet对象，然后会自动将wireApp的参数传入到ProviderSet中
	panic(wire.Build(controller.ProviderSet))
}
