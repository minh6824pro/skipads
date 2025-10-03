package skipcmd

import (
	"SkipAdsV2/config"
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/entities"
	"SkipAdsV2/redis_service"
	"context"
	"go.uber.org/zap"
)

type DatabaseInterface interface {
	CreateEventAddSkipAds(ctx context.Context, event *entities.EventAddSkipAds) error
	CreatePackage(ctx context.Context, pkg *entities.Package, games []*entities.PackageGame) error
	GetPurchasePackageByID(ctx context.Context, packageID *string) (entities.Package, error)
	GetExchangePackageByID(ctx context.Context, packageID *string) (entities.Package, error)
	ProcessEventUseSkipAds(ctx context.Context, request httpmodel.UseSkipAdsRequest) error
}

type Command struct {
	cfg    config.Config
	db     DatabaseInterface
	redis  *redis_service.RedisService
	logger *zap.Logger
}

func NewCommand(Cfg config.Config, db DatabaseInterface, redis_service *redis_service.RedisService) (*Command, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	// implement config,db
	return &Command{
		cfg:    Cfg,
		db:     db,
		redis:  redis_service,
		logger: logger,
	}, nil
}
