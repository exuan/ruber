package provider

import (
	"context"
	"time"

	rsV1 "github.com/exuan/ruber/api/recipe/service/v1"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewRecipeServiceClient(d registry.Discovery) rsV1.RecipeClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///ruber.recipe.service"),
		grpc.WithDiscovery(d),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return rsV1.NewRecipeClient(conn)
}
