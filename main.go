package main

import (
	"SkipAdsV2/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

	//TODO init repo,service,controller

	r := gin.Default()

	r.Run(":8080")

}
