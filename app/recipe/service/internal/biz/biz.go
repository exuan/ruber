package biz

import (
	"github.com/exuan/ruber/pkg/provider"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRecipe, provider.NewRedis)
