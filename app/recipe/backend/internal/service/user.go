package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/backend/v1"
)

func (bs *BackendService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return bs.b.Login(ctx, req)
}

func (bs *BackendService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return nil, nil
}

func (bs *BackendService) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	return bs.b.User(ctx, req)
}
