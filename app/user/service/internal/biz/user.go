package biz

import (
	"context"
	"fmt"

	v1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/app/user/service/internal/data/ent"
	"github.com/exuan/ruber/pkg/errors"
	"github.com/exuan/ruber/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type UserRepo interface {
	Users(ctx context.Context, req *v1.UsersRequest) (int64, []*ent.User, error)

	User(ctx context.Context, req *v1.UserRequest) (*ent.User, error)

	SaveUser(ctx context.Context, req *v1.SaveUserRequest) (string, error)

	DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (string, error)

	SetPassword(ctx context.Context, req *v1.SetPasswordRequest) (string, error)
}

type User struct {
	repo UserRepo
	log  *log.Helper
	r    *redis.Client
}

const (
	luk = "%s-%s-%s" // owner app username
)

func NewUser(repo UserRepo, logger log.Logger, r *redis.Client) *User {
	return &User{repo: repo, log: log.NewHelper(logger), r: r}
}

func (u *User) Users(ctx context.Context, req *v1.UsersRequest) (*v1.UsersReply, error) {
	total, res, err := u.repo.Users(ctx, req)

	if err != nil {
		return nil, err
	}

	items := make([]*v1.UserReply, 0)
	for _, eu := range res {
		items = append(items, &v1.UserReply{
			Id:         eu.ID,
			Username:   eu.Username,
			Phone:      eu.Phone,
			Nickname:   eu.Nickname,
			Avatar:     eu.Avatar,
			Owner:      eu.Owner,
			App:        eu.App,
			Status:     int32(eu.Status),
			Creator:    eu.Creator,
			Modifier:   eu.Modifier,
			CreateTime: eu.CreateTime,
			UpdateTime: eu.UpdateTime,
		})
	}

	return &v1.UsersReply{Total: total, Items: items}, err
}

func (u *User) User(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	res, err := u.repo.User(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		Id:         res.ID,
		Username:   res.Username,
		Phone:      res.Phone,
		Nickname:   res.Nickname,
		Avatar:     res.Avatar,
		Owner:      res.Owner,
		App:        res.App,
		Status:     int32(res.Status),
		Creator:    res.Creator,
		Modifier:   res.Modifier,
		CreateTime: res.CreateTime,
		UpdateTime: res.UpdateTime,
	}, nil
}

func (u *User) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*v1.SaveUserReply, error) {
	// Register check
	if len(req.UserId) == 0 {
		uuid := uuid.NewString()
		lk := fmt.Sprintf(luk, req.Owner, req.App, req.Username)
		if ok := util.Lock(ctx, u.r, uuid, lk, 60, 3); !ok {
			return nil, fmt.Errorf("lock failed")
		}
		defer util.UnLock(ctx, u.r, lk, uuid)

		res, err := u.repo.User(ctx, &v1.UserRequest{
			Username: req.Username,
			Owner:    req.Owner,
			App:      req.App,
		})

		if err != nil && !ent.IsNotFound(err) {
			return nil, errors.UnFound
		}

		if res != nil {
			return nil, errors.UserRegistered
		}
	}

	id, err := u.repo.SaveUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.SaveUserReply{Id: id}, nil
}

func (u *User) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserRequest, error) {
	id, err := u.repo.DeleteUser(ctx, req)
	return &v1.DeleteUserRequest{Id: id}, err
}

func (u *User) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordRequest) (*v1.VerifyPasswordReply, error) {
	res, err := u.repo.User(ctx, &v1.UserRequest{
		Username: req.Username,
		Owner:    req.Owner,
		App:      req.App,
	})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyPasswordReply{
		Id: res.ID,
		Ok: util.CheckPasswordHash(req.Password, res.Salt, res.Password),
	}, nil
}

func (u *User) SetPassword(ctx context.Context, req *v1.SetPasswordRequest) (*v1.SetPasswordReply, error) {
	id, err := u.repo.SetPassword(ctx, req)
	return &v1.SetPasswordReply{Id: id}, err
}
