package repository

import (
	"SkipAdsV2/repository/repomodel"
	"context"
	"time"
)

func (r *RepoMySQL) GetUserSkipAds(ctx context.Context, userID string) (repomodel.SkipAdsResult, error) {
	var result repomodel.SkipAdsResult
	now := time.Now()
	err := r.db.Raw(`
    SELECT 
        user_id,
        SUM(quantity - quantity_used) AS skip_ads_total
    FROM event_add_skip_ads
    WHERE user_id = ?     
      AND expires_at > ?
	  AND quantity-quantity_used>0
    GROUP BY user_id
`, userID, now).Scan(&result).Error

	if err != nil {
		return repomodel.SkipAdsResult{}, err
	}
	return result, err
}
