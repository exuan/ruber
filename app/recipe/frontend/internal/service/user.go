package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
)

func (fs *FrontendService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return fs.b.Register(ctx, req)
}

func (fs *FrontendService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return fs.b.Login(ctx, req)
}

func (fs *FrontendService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return fs.b.Logout(ctx, req)

}

func (fs *FrontendService) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	return fs.b.User(ctx, req)
}
