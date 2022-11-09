package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"

	"hello_gf/internal/model"
	"hello_gf/internal/model/entity"
	"hello_gf/internal/service"
)

type sUser struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func (s *sUser) GetAvatarUploadPath() string {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetAvatarUploadUrlPrefix() string {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) EncryptPassword(passport, password string) string {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetUserByPassportAndPassword(ctx context.Context, passport, password string) (user *entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) CheckPassportUnique(ctx context.Context, passport string) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) CheckNicknameUnique(ctx context.Context, nickname string) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) Register(ctx context.Context, in model.UserRegisterInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) UpdatePassword(ctx context.Context, in model.UserPasswordInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetProfileById(ctx context.Context, userId uint) (out *model.UserGetProfileOutput, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetProfile(ctx context.Context) (*model.UserGetProfileOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) UpdateAvatar(ctx context.Context, in model.UserUpdateAvatarInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) UpdateProfile(ctx context.Context, in model.UserUpdateProfileInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) Disable(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetList(ctx context.Context, in model.UserGetContentListInput) (out *model.UserGetListOutput, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetMessageList(ctx context.Context, in model.UserGetMessageListInput) (out *model.UserGetMessageListOutput, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) GetUserStats(ctx context.Context, userId uint) (map[string]int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) IsCtxAdmin(ctx context.Context) bool {
	//TODO implement me
	panic("implement me")
}

func (s *sUser) IsAdmin(ctx context.Context, userId uint) bool {
	//TODO implement me
	panic("implement me")
}

func init() {
	user := New()
	// 启动时创建头像存储目录
	if !gfile.Exists(user.avatarUploadPath) {
		if err := gfile.Mkdir(user.avatarUploadPath); err != nil {
			g.Log().Fatal(gctx.New(), err)
		}
	}
	service.RegisterUser(user)
}

func New() *sUser {
	return &sUser{
		avatarUploadPath:      g.Cfg().MustGet(gctx.New(), `upload.path`).String() + `/avatar`,
		avatarUploadUrlPrefix: `/upload/avatar`,
	}
}
