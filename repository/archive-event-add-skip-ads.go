package repository

import (
	"SkipAdsV2/entities"
	"time"
)

func (r *RepoMySQL) ArchiveEventAddSkipAds() error {
	batchSize := 1000
	for {
		var ids []int64

		// get batch id
		if err := r.db.
			Model(&entities.EventAddSkipAds{}).
			Where("quantity = quantity_used OR  expires_at < ?", time.Now()).
			Limit(batchSize).
			Pluck("id", &ids).Error; err != nil {
			return err
		}

		if len(ids) == 0 {
			break
		}

		tx := r.db.Begin()
		if tx.Error != nil {
			return tx.Error
		}

		// insert into archive
		if err := tx.Exec(`
			INSERT INTO event_add_skip_ads_archives
			SELECT *
			FROM event_add_skip_ads
			WHERE id IN (?)`, ids).Error; err != nil {
			tx.Rollback()
			return err
		}

		// delete in original table
		if err := tx.Where("id IN (?)", ids).
			Delete(&entities.EventAddSkipAds{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}

		// not enough batch then no data left
		if len(ids) < batchSize {
			break
		}
	}
	return nil

}
