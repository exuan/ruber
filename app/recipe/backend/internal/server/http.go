package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	v1 "github.com/exuan/ruber/api/recipe/backend/v1"
	"github.com/exuan/ruber/app/recipe/backend/internal/conf"
	"github.com/exuan/ruber/app/recipe/backend/internal/service"
	"github.com/exuan/ruber/pkg/util/format"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	mm "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
)

func NewHTTPServer(c *conf.Server, bs *service.BackendService, r *redis.Client, logger log.Logger) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			mm.Server(mm.WithPropagatedPrefix("x-ruber-")),
			recovery.Recovery(),
			Auth(r),
		),

		khttp.RequestDecoder(func(r *http.Request, i interface{}) error {
			return binding.Default(r.Method, format.HeaderFilterFlags(r.Header.Get("Content-Type"))).Bind(r, i)
		}),

		khttp.ResponseEncoder(func(w http.ResponseWriter, request *http.Request, i interface{}) error {
			if i == nil {
				i = map[int]int{}
			}

			res, err := jsoniter.Marshal(format.Output{
				Data: i,
			})
			if err != nil {
				return err
			}
			w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json; charset=utf-8")
			_, _ = w.Write(res)
			return nil
		}),

		khttp.ErrorEncoder(func(w http.ResponseWriter, request *http.Request, err error) {
			res, err := jsoniter.Marshal(format.Output{
				1,
				err.Error(),
				format.DefaultData{},
			})

			w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json; charset=utf-8")
			_, _ = w.Write(res)
		}),
	}

	if c.Http.Network != "" {
		opts = append(opts, khttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, khttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, khttp.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := khttp.NewServer(opts...)
	v1.RegisterBackendHTTPServer(srv, bs)
	return srv
}

func Auth(r *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.Unauthorized("", "3")
			}
			ht, ok := tr.(*khttp.Transport)
			if !ok {
				return nil, errors.Unauthorized("", "2")
			}

			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return nil, errors.Unauthorized("", "1")
			}
			token := md.Get("x-ruber-backend-token")
			if ht.Request().RequestURI != "/v1/login" {
				fmt.Println(token)
				if len(token) == 0 {
					return nil, errors.Unauthorized("", "5")
				}
				id, err := r.Get(ctx, token).Result()
				fmt.Println(id, err)
				if err != nil {
					return nil, errors.Unauthorized("", "")
				}

				if len(id) == 0 {
					return nil, errors.Unauthorized("", "")
				}

				md.Set("user_id", id)
				r.Expire(ctx, token, 24*time.Hour)
			}
			return handler(ctx, req)
		}
	}
}
