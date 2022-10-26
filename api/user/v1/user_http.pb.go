package v1

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	_SignInRequestValidateMap = map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	_UserProfileRequestValidateMap = map[string]string{
		"name": "required|length:6,16",
	}
)

func RegisterUserHTTPServer(s *ghttp.Server, srv UserHttpServer) {
	// 这里的 /user 根据 proto 中定义的服务自动生成
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.POST("/sign_in", _userProfileHttpHandler(srv))
		group.GET("/user_profile", _userSignInHttpHandler(srv))
	})
}

func _userProfileHttpHandler(srv UserHttpServer) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		var (
			in  *UserProfileRequest
			err error
			ctx context.Context
		)
		// 这里需要让用户可以自定义错误代码
		if err = r.Parse(&in); err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		if err = g.Validator().Rules(_UserProfileRequestValidateMap).Data(in).Run(ctx); err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		// 这里可以让用户自定义注入的ctx
		profile, err := srv.UserProfile(ctx, in)
		if err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		_ = r.Response.WriteJson(profile)
	}
}

func _userSignInHttpHandler(srv UserHttpServer) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		var (
			in  *SignInRequest
			err error
			ctx context.Context
		)
		// 这里需要让用户可以自定义错误代码
		if err = r.Parse(&in); err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		if err = g.Validator().Rules(_SignInRequestValidateMap).Data(in).Run(ctx); err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		// 这里可以让用户自定义注入的ctx
		profile, err := srv.UserSignIn(ctx, in)
		if err != nil {
			// 统一的返回层，可以搭配gerror
			return
		}
		_ = r.Response.WriteJson(profile)
	}
}

type UserHttpServer interface {
	UserSignIn(ctx context.Context, in *SignInRequest) (*UserSignResponse, error)
	UserProfile(ctx context.Context, in *UserProfileRequest) (*UserProfileResponse, error)
}
