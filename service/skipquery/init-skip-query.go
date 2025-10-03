package skipquery

import (
	"SkipAdsV2/config"
	"SkipAdsV2/repository/repomodel"
	"context"
	"go.uber.org/zap"
)

type DatabaseInterface interface {
	GetUserSkipAds(ctx context.Context, userID string) (repomodel.SkipAdsResult, error)
}

type Query struct {
	cfg    config.Config
	db     DatabaseInterface
	logger *zap.Logger
}

func NewQuery(cfg config.Config, db DatabaseInterface) (*Query, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	// implement config, db
	return &Query{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}, nil
}
