package user

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"google.golang.org/grpc"
)

// 这个方法需要到时候封装到gf里面
type gogfApp struct {
	ghc *ghttp.Server
	grc *grpc.Server
	log *glog.Logger
}

// 这个方法需要到时候封装到gf里面
func (*gogfApp) Run() error {
	return nil
}

func main() {
	// 最前面部分预留做参数解析
	// =========
	// 这部分相当于是模板的文件，用户不需要手动修改，全自动生成
	app, cleanUp, err := wireApp()
	if err != nil {
		panic(err.Error())
	}
	defer cleanUp()
	err = app.Run()
	if err != nil {
		panic(err.Error())
	}
}

func newApp(logger *glog.Logger, gs *grpc.Server, hs *ghttp.Server) *gogfApp {
	return &gogfApp{
		ghc: hs,
		grc: gs,
		log: logger,
	}
}
