package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"

	v1 "hello_gf/api/generated/http/echo/v1"
	"hello_gf/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start focus server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			//oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// 静态目录设置
			//uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			//if uploadPath == "" {
			//	g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			//}
			//s.AddStaticPath("/upload", uploadPath)

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
				//service.Middleware().Ctx,
				//service.Middleware().ResponseHandler,
				)
				group.Bind(v1.NewEchoApi(controller.User))

			})
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `Focus Project`
	openapi.Info.Description = ``

}
