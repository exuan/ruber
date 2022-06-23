package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	v1 "github.com/exuan/ruber/api/recipe/frontend/v1"
	"github.com/exuan/ruber/app/recipe/frontend/internal/conf"
	"github.com/exuan/ruber/app/recipe/frontend/internal/service"
	"github.com/exuan/ruber/pkg/errors"
	"github.com/exuan/ruber/pkg/middleware/bind"
	"github.com/exuan/ruber/pkg/util/format"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	mm "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	jsoniter "github.com/json-iterator/go"
)

func NewHTTPServer(c *conf.Server, fs *service.FrontendService, r *redis.Client, logger log.Logger) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			mm.Server(mm.WithPropagatedPrefix("x-ruber-")),
			recovery.Recovery(),
			Auth(r),
			bind.Bind(),
		),

		khttp.RequestDecoder(func(r *http.Request, i interface{}) error {
			if err := binding.Default(r.Method, format.HeaderFilterFlags(r.Header.Get("Content-Type"))).Bind(r, i); err != nil {
				fmt.Println(err)
				return errors.ParameterErrors
			}
			return nil
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
			se, ok := err.(*errors.Error)
			if !ok {
				fmt.Println(se)
				se = errors.InternalServerError
			}

			res, _ := jsoniter.Marshal(format.Output{
				Code: se.Code,
				Msg:  se.Msg,
				Data: format.DefaultData{},
			})

			w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json; charset=utf-8")
			_, _ = w.Write(res)
		}),
		khttp.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	v1.RegisterFrontendHTTPServer(srv, fs)
	return srv
}

func Auth(r *redis.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.ServiceUnavailable
			}
			ht, ok := tr.(*khttp.Transport)
			if !ok {
				return nil, errors.ServiceUnavailable
			}

			md, ok := metadata.FromServerContext(ctx)
			if !ok {
				return nil, errors.ServiceUnavailable
			}

			if token := md.Get("x-ruber-token"); len(token) > 0 {
				if id, err := r.Get(ctx, token).Result(); err == nil {
					md.Set("user_id", id)
					r.Expire(ctx, token, 24*time.Hour)
				}
			}
			if ht.Request().URL.Path != "/v1/login" &&
				ht.Request().URL.Path != "/v1/register" &&
				ht.Request().URL.Path != "/v1/recipe/categories" &&
				ht.Request().URL.Path != "/v1/recipe" &&
				ht.Request().URL.Path != "/v1/recipes" {

				if len(md.Get("user_id")) == 0 {
					return nil, errors.Unauthorized
				}

			}

			return handler(ctx, req)
		}
	}
}
