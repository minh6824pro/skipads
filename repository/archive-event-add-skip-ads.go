package repository

import (
	"SkipAdsV2/entities"
	"database/sql"
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

		// By default, MySQL uses REPEATABLE READ isolation.
		// Under REPEATABLE READ, SELECT ... FOR UPDATE will keep locks on rows
		// even if they no longer match the query condition later in the transaction.
		// This can cause deadlocks between two transactions:
		//   - "consume skip ads" locks usable rows (quantity > quantity_used, not expired)
		//   - "archive" locks expired or fully used rows
		// Because REPEATABLE READ holds locks more aggressively, both transactions
		// may wait on each other’s locks, resulting in a deadlock.

		// Set the isolation level for this transaction to READ COMMITTED.
		// This helps prevent deadlocks when archiving and consuming skip ads concurrently.
		// Reason: consuming skip ads locks only usable events (not expired and quantity > quantity_used),
		// while archiving locks expired or fully used events. Using READ COMMITTED avoids unnecessary conflicts.
		tx := r.db.Begin(&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
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
