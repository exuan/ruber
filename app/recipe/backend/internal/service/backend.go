package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/backend/v1"
	"github.com/exuan/ruber/app/recipe/backend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type BackendService struct {
	v1.UnimplementedBackendServer

	b   *biz.Backend
	log *log.Helper
}

func NewBackendService(b *biz.Backend, logger log.Logger) *BackendService {
	return &BackendService{b: b, log: log.NewHelper(logger)}
}

func (bs *BackendService) UploadUrl(ctx context.Context, req *v1.UploadUrlRequest) (*v1.UploadUrlReply, error) {
	return bs.b.UploadUrl(ctx, req)
}
