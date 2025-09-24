package repository

import (
	"SkipAdsV2/entities"
	"context"
)

func (r *RepoMySQL) CreateEventAddSkipAds(ctx context.Context, event *entities.EventAddSkipAds) error {
	err := r.db.Create(&event).Error
	if err != nil {
		return err
	}
	return nil
}
