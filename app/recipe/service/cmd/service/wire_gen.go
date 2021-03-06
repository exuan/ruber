// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/exuan/ruber/app/recipe/service/internal/biz"
	"github.com/exuan/ruber/app/recipe/service/internal/conf"
	"github.com/exuan/ruber/app/recipe/service/internal/data"
	"github.com/exuan/ruber/app/recipe/service/internal/server"
	"github.com/exuan/ruber/app/recipe/service/internal/service"
	"github.com/exuan/ruber/pkg/provider"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, string2 string, redisConfig *provider.RedisConfig, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	recipeRepo := data.NewRecipeRepo(dataData, logger)
	client, cleanup2, err := provider.NewRedis(redisConfig, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	recipe := biz.NewRecipe(recipeRepo, logger, client)
	recipeService := service.NewRecipeService(recipe, logger)
	httpServer := server.NewHTTPServer(confServer, recipeService, logger)
	grpcServer := server.NewGRPCServer(confServer, recipeService, logger)
	registrar := provider.NewRegistrar(string2)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
