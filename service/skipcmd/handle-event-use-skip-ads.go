package skipcmd

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"context"
)

func (cmd *Command) HandleEventUseSkipAds(ctx context.Context, request httpmodel.UseSkipAdsRequest) error {

	// check redis health
	if cmd.redis.IsAlive.Load() {
		// Get redis Lock
		token, err := cmd.CreateRedisLock(ctx, request.UserID)
		if err != nil {
			cmd.logger.Error(err.Error())
		} else {
			defer cmd.redis.ReleaseLock(ctx, request.UserID, token)
		}
	}

	// Process with DB
	return cmd.db.ProcessEventUseSkipAds(ctx, request)
}
