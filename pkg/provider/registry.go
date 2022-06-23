package provider

import (
	"time"

	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewRegistrar(addr string) registry.Registrar {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	r := etcd.New(cli)
	return r
}

func NewDiscovery(addr string) registry.Discovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	return etcd.New(cli)
}
