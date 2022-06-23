package biz

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/backend/v1"
	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
	usV1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/exuan/ruber/pkg/provider"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type Backend struct {
	rc  rsV1.RecipeClient
	uc  usV1.UserClient
	log *log.Helper
	r   *redis.Client
	a   provider.Attachment
}

func NewBackend(rc rsV1.RecipeClient, uc usV1.UserClient, logger log.Logger, r *redis.Client, a provider.Attachment) *Backend {
	return &Backend{
		rc:  rc,
		uc:  uc,
		log: log.NewHelper(logger),
		r:   r,
		a:   a,
	}
}

func (b *Backend) UploadUrl(_ context.Context, req *v1.UploadUrlRequest) (*v1.UploadUrlReply, error) {
	pu, u, err := b.a.PresignedURL("", req.Ext)
	return &v1.UploadUrlReply{
		PresignedUrl: pu,
		Url:          u,
	}, err
}
