package biz

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
	usV1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/pkg/provider"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type Frontend struct {
	rc  rsV1.RecipeClient
	uc  usV1.UserClient
	log *log.Helper
	r   *redis.Client
	a   provider.Attachment
}

func NewFrontend(rc rsV1.RecipeClient, uc usV1.UserClient, logger log.Logger, r *redis.Client, a provider.Attachment) *Frontend {
	return &Frontend{
		rc:  rc,
		uc:  uc,
		log: log.NewHelper(logger),
		r:   r,
		a:   a,
	}
}

func (f *Frontend) Index(ctx context.Context, req *v1.IndexRequest) (*v1.IndexReply, error) {
	res, err := f.rc.Index(ctx, &rsV1.IndexRequest{})
	if err != nil {
		return nil, err
	}

	return &v1.IndexReply{Carousel: res.Carousel}, nil
}
