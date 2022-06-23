package biz

import (
	"context"
	"time"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
	usV1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/pkg/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/google/uuid"
)

const (
	owner = "recipe"
	app   = "frontend"
)

func (f *Frontend) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	_, err := f.uc.SaveUser(ctx, &usV1.SaveUserRequest{
		Username: req.Username,
		Nickname: req.Username,
		Password: req.Password,
		Owner:    owner,
		App:      app,
	})

	if err != nil {
		return nil, errors.UserRegistered
	}
	return &v1.RegisterReply{}, err
}

func (f *Frontend) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	reply, err := f.uc.VerifyPassword(ctx, &usV1.VerifyPasswordRequest{
		Username: req.Username,
		Password: req.Password,
		Owner:    owner,
		App:      app,
	})

	if err != nil || !reply.Ok {
		return nil, errors.IncorrectPassword
	}

	token := uuid.NewString()
	if err = f.r.SetEX(ctx, token, reply.Id, 24*time.Hour).Err(); err != nil {
		return nil, errors.InternalServerError
	}

	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (f *Frontend) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.Unauthorized
	}

	token := md.Get("x-ruber-token")

	if len(token) == 0 {
		return nil, errors.Unauthorized
	}
	_, err := f.r.Del(ctx, token).Result()

	if err != nil {
		return nil, errors.ServiceUnavailable
	}
	return &v1.LogoutReply{}, nil
}

func (f *Frontend) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.Unauthorized
	}

	res, err := f.uc.User(ctx, &usV1.UserRequest{
		Id:    md.Get("user_id"),
		Owner: owner,
		App:   app,
	})

	if err != nil {
		return nil, errors.UnFound
	}

	return &v1.UserReply{
		Id:         res.Id,
		Username:   res.Username,
		Nickname:   res.Nickname,
		Phone:      res.Phone,
		Avatar:     res.Avatar,
		CreateTime: res.CreateTime,
		UpdateTime: res.UpdateTime,
	}, nil
}
