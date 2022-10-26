package service

import (
	"context"

	v1 "hello_gf/api/user/v1"
)

type User struct {
	v1.UnimplementedUserServer
}

// 这部分相当于薄薄的一层，不建议用户逻辑写在这里，这里只是做一些参数的转换和错误判断
// 层级可能还是需要往上提一层，这也是在用Kratos时发现的问题
// 这里并不算是真正的服务层，这里只是连接自动生成和用户自己业务逻辑的地方
func (u *User) UserSignIn(ctx context.Context, in *v1.SignInRequest) (*v1.UserSignResponse, error) {
	return nil, nil
}

func (u *User) UserProfile(ctx context.Context, in *v1.UserProfileRequest) (*v1.UserProfileResponse, error) {
	return nil, nil
}
