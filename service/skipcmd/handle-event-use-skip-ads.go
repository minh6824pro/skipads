package skipcmd

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"context"
)

func (cmd *Command) HandleEventUseSkipAds(ctx context.Context, request httpmodel.UseSkipAdsRequest) error {

	// Get Redis Lock

	// Process with DB
	return cmd.db.ProcessEventUseSkipAds(ctx, request)
}
