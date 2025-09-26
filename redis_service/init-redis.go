package redis_service

import (
	"SkipAdsV2/config"
	"context"
	"github.com/redis/go-redis/v9"
	"sync/atomic"
)

const (
	UserSkipAdsLockKeyPattern = "user-skip-ads-lock-%d"
)

type RedisService struct {
	RedisClient *redis.Client
	IsAlive     atomic.Bool
}

func NewRedis(Cfg config.Config) (*RedisService, error) {
	addr := Cfg.Redis.URI
	password := Cfg.Redis.Password
	db := Cfg.Redis.Db

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// check redis connection
	err := redisClient.Ping(context.Background()).Err()
	redisService := &RedisService{RedisClient: redisClient}
	if err != nil {
		redisService.IsAlive.Store(false)

	} else {
		redisService.IsAlive.Store(true)
	}
	return redisService, err
}
