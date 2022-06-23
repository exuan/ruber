package data

import (
	"context"

	v1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/app/user/service/internal/biz"
	"github.com/exuan/ruber/app/user/service/internal/data/ent"
	"github.com/exuan/ruber/app/user/service/internal/data/ent/user"
	"github.com/exuan/ruber/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) Users(ctx context.Context, req *v1.UsersRequest) (int64, []*ent.User, error) {
	sess := ur.data.db.User.Query().
		Where(user.StatusNEQ(-1))

	if len(req.Id) > 0 {
		sess.Where(user.ID(req.Id))
	}
	if len(req.Username) > 0 {
		sess.Where(user.Username(req.Username))
	}
	if len(req.Phone) > 0 {
		sess.Where(user.Phone(req.Phone))
	}
	if len(req.Nickname) > 0 {
		sess.Where(user.ID(req.Nickname))
	}
	if len(req.Owner) > 0 {
		sess.Where(user.Owner(req.Owner))
	}
	if len(req.App) > 0 {
		sess.Where(user.App(req.App))
	}

	total, err := sess.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	res, err := sess.Order(ent.Desc(user.FieldCreateTime)).
		Limit(int(req.Size)).
		Offset(int((req.Page - 1) * req.Size)).
		All(ctx)

	return int64(total), res, err
}

func (ur *userRepo) User(ctx context.Context, req *v1.UserRequest) (*ent.User, error) {
	sess := ur.data.db.Debug().User.Query().
		Where(user.StatusNEQ(-1))

	if len(req.Id) > 0 {
		sess.Where(user.ID(req.Id))
	}
	if len(req.Username) > 0 {
		sess.Where(user.Username(req.Username))
	}
	if len(req.Phone) > 0 {
		sess.Where(user.Phone(req.Phone))
	}
	if len(req.Nickname) > 0 {
		sess.Where(user.Nickname(req.Nickname))
	}
	if len(req.Owner) > 0 {
		sess.Where(user.Owner(req.Owner))
	}
	if len(req.App) > 0 {
		sess.Where(user.App(req.App))
	}

	return sess.Only(ctx)
}

func (ur *userRepo) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (string, error) {
	var err error
	var eu *ent.User

	if len(req.Id) > 0 {
		eu, err = ur.data.db.User.UpdateOneID(req.Id).
			SetUsername(req.Username).
			SetPhone(req.Phone).
			SetNickname(req.Nickname).
			SetAvatar(req.Avatar).
			SetOwner(req.Owner).
			SetApp(req.App).
			SetModifier(req.UserId).
			Save(ctx)

	} else {
		increment, _ := uuid.NewRandom()
		slat := increment.String()[:8]
		password, _ := util.HashPassword(req.Password, slat)

		eu, err = ur.data.db.User.Create().
			SetID(uuid.NewString()).
			SetUsername(req.Username).
			SetPhone(req.Phone).
			SetPassword(password).
			SetSalt(slat).
			SetNickname(req.Nickname).
			SetAvatar(req.Avatar).
			SetOwner(req.Owner).
			SetApp(req.App).
			SetStatus(0).
			SetModifier(req.UserId).
			SetCreator(req.UserId).
			Save(ctx)
	}

	if err != nil {
		return "", err
	}
	return eu.ID, err
}

func (ur *userRepo) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (string, error) {
	et, err := ur.data.db.User.
		UpdateOneID(req.Id).
		SetModifier(req.UserId).
		SetStatus(-1).
		Save(ctx)

	return et.ID, err
}

func (ur *userRepo) SetPassword(ctx context.Context, req *v1.SetPasswordRequest) (string, error) {
	increment, _ := uuid.NewRandom()
	slat := increment.String()[:8]
	password, _ := util.HashPassword(req.Password, slat)

	et, err := ur.data.db.User.
		UpdateOneID(req.Id).
		SetPassword(password).
		SetSalt(slat).
		Save(ctx)

	return et.ID, err
}
