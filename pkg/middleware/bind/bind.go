package bind

import (
	"context"
	"log"

	"github.com/exuan/ruber/pkg/errors"
	"github.com/exuan/ruber/pkg/util/format"
	"github.com/gin-gonic/gin/binding"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// Bind 此处略妖
func Bind() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.ParameterErrors
			}
			ht, ok := tr.(*khttp.Transport)
			if !ok {
				return nil, errors.ParameterErrors
			}

			if format.HeaderFilterFlags(ht.Request().Header.Get("Content-Type")) != binding.MIMEJSON {
				if err := binding.Default(
					ht.Request().Method,
					format.HeaderFilterFlags(ht.Request().Header.Get("Content-Type"))).Bind(ht.Request(),
					req); err != nil {
					log.Println(err, req, format.HeaderFilterFlags(ht.Request().Header.Get("Content-Type")))
					return nil, errors.ParameterErrors
				}
			}

			return handler(ctx, req)
		}
	}
}
