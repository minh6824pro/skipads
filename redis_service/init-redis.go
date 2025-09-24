package redis_service

import (
	"SkipAdsV2/config"
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	RedisClient *redis.Client
}

func NewRedis(Cfg config.Config) (*RedisService, error) {
	addr := Cfg.Redis.URI
	password := Cfg.Redis.Password
	dbStr := Cfg.Redis.DbString

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbStr,
	})

	// Kiểm tra kết nối Redis
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	} else {
		return nil, err
	}
	return &RedisService{RedisClient: redisClient}, nil
}
