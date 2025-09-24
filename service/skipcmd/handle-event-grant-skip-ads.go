package skipcmd

import (
	"SkipAdsV2/entities"
	"context"
)

func (cmd *Command) HandleEventGrantSkipAds(ctx context.Context, eventExchange *entities.EventAddSkipAds) error {
	// TODO check user exists

	eventExchange.QuantityUsed = 0
	eventExchange.Type = entities.EventAddSkipAdsGrant

	// create event
	err := cmd.db.CreateEventAddSkipAds(ctx, eventExchange)
	if err != nil {
		return err
	}
	return nil

}
