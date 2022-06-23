package provider

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type (
	RedisConfig struct {
		Addr         string `json:"addr"`
		Password     string `json:"password"`
		DB           int    `json:"db"`
		ReadTimeout  int64  `json:"read_timeout"`
		WriteTimeout int64  `json:"write_timeout"`
	}
)

func NewRedis(conf *RedisConfig, logger log.Logger) (*redis.Client, func(), error) {
	log := log.NewHelper(logger)

	client := redis.NewClient(
		&redis.Options{
			Addr:         conf.Addr,
			Password:     conf.Password,
			DB:           conf.DB,
			ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
		})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Errorf("failed ping to redis: %v", err)
		return nil, nil, err
	}

	return client, func() {
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
