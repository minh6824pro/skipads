package main

import (
	"SkipAdsV2/config"
	"SkipAdsV2/controller/userskipadshttp"
	"SkipAdsV2/redis_service"
	"SkipAdsV2/repository"
	"SkipAdsV2/service/skipcmd"
	"SkipAdsV2/service/skipquery"
	"context"
	"go.uber.org/zap"
	"time"
)

func main() {

	// Init Logger
	lg, _ := zap.NewProduction()
	defer lg.Sync()

	cfg, err := config.NewConfig()
	if err != nil {
		lg.Panic("cannot init config", zap.Error(err))
	}
	lg.Info("config", zap.Any("config", cfg))

	repo, err := repository.NewRepoMysql(cfg)
	if err != nil {
		lg.Panic("cannot connect to mysql", zap.Error(err))
	}
	err = repo.InitTable()
	if err != nil {
		lg.Panic("InitTable failed", zap.Error(err))
	}

	redis, err := redis_service.NewRedis(cfg)
	if err != nil {
		lg.Info("cannot init redis", zap.Error(err))
	}
	command, err := skipcmd.NewCommand(cfg, repo, redis)
	if err != nil {
		lg.Panic("cannot init commands skip ", zap.Error(err))
	}

	query, err := skipquery.NewQuery(cfg, repo)
	if err != nil {
		lg.Panic("cannot init query skip ", zap.Error(err))
	}

	ginHttp, err := userskipadshttp.NewHttpServer(cfg, command, query)
	if err != nil {
		lg.Panic("cannot init http server user skip ads ", zap.Error(err))
	}

	redis.StartRedisHealthCheck(context.Background(), 5*time.Second)

	// Seed  user data
	//repo.SeedUser()

	// Seed purchase,exchange
	//repo.SeedSkipAds()

	// cron job clean not usable event add skip ads
	//go func() {
	//	time.Sleep(2 * time.Second) //
	//	if err := repo.ArchiveEventAddSkipAds(); err != nil {
	//		log.Println("[startup] Archive failed:", err)
	//	}
	//}()

	ginHttp.StartWithGracefulShutdown()

}
