package repository

import (
	"SkipAdsV2/entities"
	"context"
	"database/sql"
	"time"
)

func (r *RepoMySQL) ArchiveEventAddSkipAds(ctx context.Context) error {
	batchSize := 1000

	for {
		// Start transaction with READ COMMITTED isolation level
		// This avoids holding locks on rows that no longer match the query condition,
		// reducing the chance of deadlocks when multiple workers run concurrently.
		tx := r.db.WithContext(ctx).Begin(&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})

		if tx.Error != nil {
			return tx.Error
		}

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		// Select a batch of IDs with row locking
		var ids []int64
		if err := tx.Raw(`
			SELECT id
			FROM event_add_skip_ads
			WHERE quantity = quantity_used OR expires_at < ?
			ORDER BY id
			LIMIT ?
		`, time.Now(), batchSize).Pluck("id", &ids).Error; err != nil {
			tx.Rollback()
			return err
		}

		// No more rows to archive
		if len(ids) == 0 {
			tx.Rollback()
			break
		}

		// Archive records by copying them into the archive table
		if err := tx.Exec(`
			INSERT INTO event_add_skip_ads_archives
			SELECT *
			FROM event_add_skip_ads
			WHERE id IN (?)
		`, ids).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Delete the archived records from the original table
		if err := tx.Where("id IN (?)", ids).
			Delete(&entities.EventAddSkipAds{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Commit transaction
		if err := tx.Commit().Error; err != nil {
			return err
		}

		// If less than batchSize was processed, it means no data left
		if len(ids) == 0 {
			break
		}

	}

	return nil
}
