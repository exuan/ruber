package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	v1.UnimplementedUserServer

	b   *biz.User
	log *log.Helper
}

func NewUserService(b *biz.User, logger log.Logger) *UserService {
	return &UserService{b: b, log: log.NewHelper(logger)}
}

func (us *UserService) Users(ctx context.Context, req *v1.UsersRequest) (*v1.UsersReply, error) {
	return us.b.Users(ctx, req)
}

func (us *UserService) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	return us.b.User(ctx, req)
}

func (us *UserService) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*v1.SaveUserReply, error) {
	return us.b.SaveUser(ctx, req)
}

func (us *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserRequest, error) {
	return us.b.DeleteUser(ctx, req)
}

func (us *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordReply, error) {
	return us.b.VerifyPassword(ctx, req)
}

func (us *UserService) SetPassword(ctx context.Context, req *v1.SetPasswordRequest) (*v1.SetPasswordReply, error) {
	return us.b.SetPassword(ctx, req)
}
