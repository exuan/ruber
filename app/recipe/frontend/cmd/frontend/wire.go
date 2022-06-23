//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/exuan/ruber/app/recipe/frontend/internal/biz"
	"github.com/exuan/ruber/app/recipe/frontend/internal/conf"
	"github.com/exuan/ruber/app/recipe/frontend/internal/server"
	"github.com/exuan/ruber/app/recipe/frontend/internal/service"
	"github.com/exuan/ruber/pkg/provider"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Server, string, *provider.RedisConfig, *provider.AttachmentConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
