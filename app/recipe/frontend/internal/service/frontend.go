package service

import (
	"context"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
	"github.com/exuan/ruber/app/recipe/frontend/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FrontendService struct {
	v1.UnimplementedFrontendServer

	b   *biz.Frontend
	log *log.Helper
}

func NewFrontendService(b *biz.Frontend, logger log.Logger) *FrontendService {
	return &FrontendService{b: b, log: log.NewHelper(logger)}
}

func (fs *FrontendService) Index(ctx context.Context, req *v1.IndexRequest) (*v1.IndexReply, error) {
	return fs.b.Index(ctx, req)
}
