package provider

import (
	"context"
	"time"

	usV1 "github.com/exuan/ruber/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewUserServiceClient(d registry.Discovery) usV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///ruber.user.service"),
		grpc.WithDiscovery(d),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return usV1.NewUserClient(conn)
}
