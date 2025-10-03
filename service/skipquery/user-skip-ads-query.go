package skipquery

import (
	"SkipAdsV2/repository/repomodel"
	"context"
)

func (q *Query) GetUserSkipAds(ctx context.Context, userID string) (repomodel.SkipAdsResult, error) {
	userSkipAds, err := q.db.GetUserSkipAds(ctx, userID)
	if err != nil {
		return userSkipAds, err
	}
	if userSkipAds.UserID != userID {
		userSkipAds.UserID = userID
	}
	return userSkipAds, nil
}
