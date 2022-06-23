package biz

import (
	"github.com/exuan/ruber/pkg/provider"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewBackend,
	provider.NewDiscovery,
	provider.NewRecipeServiceClient,
	provider.NewUserServiceClient,
	provider.NewRedis,
	provider.NewAttachment,
)
