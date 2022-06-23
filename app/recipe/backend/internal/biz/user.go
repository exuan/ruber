package biz

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/exuan/ruber/api/recipe/backend/v1"
	usV1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/google/uuid"
)

const (
	owner = "recipe"
	app   = "backend"
)

func (b *Backend) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	reply, err := b.uc.VerifyPassword(ctx, &usV1.VerifyPasswordRequest{
		Username: req.Username,
		Password: req.Password,
		Owner:    owner,
		App:      app,
	})
	fmt.Println(err, reply)
	if err != nil || !reply.Ok {
		return nil, v1.ErrorIncorrectPassword("incorrect password")
	}

	token := uuid.NewString()
	if err = b.r.SetEX(ctx, token, reply.Id, 24*time.Hour).Err(); err != nil {
		return nil, v1.ErrorInternalServerError("server error")
	}

	return &v1.LoginReply{
		Token: token,
	}, nil
}

func (b *Backend) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("", "")
	}
	token := md.Get("x-ruber-backend-token")

	if len(token) == 0 {
		return nil, errors.Unauthorized("", "")
	}
	_, _ = b.r.Del(ctx, token).Result()
	return &v1.LogoutReply{}, nil
}

func (b *Backend) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil, errors.Unauthorized("", "")
	}
	id := md.Get("id")

	res, err := b.uc.User(ctx, &usV1.UserRequest{
		Id:    id,
		Owner: owner,
		App:   app,
	})

	if err != nil {
		return nil, errors.NotFound("", "")
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
