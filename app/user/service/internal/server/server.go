package server

import (
	"github.com/exuan/ruber/pkg/provider"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, provider.NewRegistrar)
