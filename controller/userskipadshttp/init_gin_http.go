package userskipadshttp

import (
	"SkipAdsV2/config"
	"SkipAdsV2/service/skipcmd"
	"SkipAdsV2/service/skipquery"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	requestid "github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"

	"time"
)

// GinHttp server struct
type GinHttp struct {
	engine      *gin.Engine
	cfg         config.Config
	logger      *zap.Logger
	command     *skipcmd.Command
	query       *skipquery.Query
	healthCheck bool
}

// NewHttpServer tạo Gin server
func NewHttpServer(cfg config.Config, cmd *skipcmd.Command, query *skipquery.Query) (*GinHttp, error) {
	logger, _ := zap.NewProduction()
	// Lưu ý: không defer logger.Sync() ở đây nếu muốn dùng lâu dài
	// defer logger.Sync()

	g := &GinHttp{
		engine:      gin.New(),
		cfg:         cfg,
		logger:      logger,
		command:     cmd,
		query:       query,
		healthCheck: true,
	}

	g.initRouters()
	return g, nil
}

// initRouters thiết lập routes
func (g *GinHttp) initRouters() {
	r := g.engine

	// Request ID (đầu tiên)
	r.Use(requestid.New())

	// Recovery (thứ 2 - catch panics)
	r.Use(gin.Recovery())

	// Logging (cuối cùng - log mọi thứ)
	r.Use(g.LoggingRequest())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	// Health check
	if g.healthCheck {
		r.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	// Example routes
	r.POST("/purchases", g.HandlePurchasePackage)
	r.GET("/skipads/:user_id", g.HandleGetUserSkipAds)
	r.POST("/exchanges", g.HandleExchangePackage)
	r.POST("/grants", g.HandleGrantSkipAds)
	r.POST("/skipads", g.HandleUseSkipAds)
}

func (g *GinHttp) StartWithGracefulShutdown() {
	engine := g.engine
	lg := g.logger
	cfg := g.cfg

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Http.Port),
		Handler: engine,
	}

	// start server trong goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lg.Panic("can't start Gin server", zap.Error(err))
		}
	}()
	lg.Info(fmt.Sprintf("Service %s started on port %d", cfg.ServiceName, cfg.Http.Port))

	// handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // Ctrl+C
	<-quit
	lg.Info(fmt.Sprintf("shutting down %s...", cfg.ServiceName))

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		lg.Fatal(fmt.Sprintf("force shutdown %s", cfg.ServiceName))
	}

	lg.Info(fmt.Sprintf("server %s stopped gracefully", cfg.ServiceName))
}
