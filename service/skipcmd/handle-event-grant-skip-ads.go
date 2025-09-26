package skipcmd

import (
	"SkipAdsV2/entities"
	"context"
)

func (cmd *Command) HandleEventGrantSkipAds(ctx context.Context, eventGrant *entities.EventAddSkipAds) error {
	// TODO check user exists

	eventGrant.QuantityUsed = 0
	eventGrant.Type = entities.EventAddSkipAdsGrant
	eventGrant.SetPriority()

	// create event
	err := cmd.db.CreateEventAddSkipAds(ctx, eventGrant)
	if err != nil {
		return err
	}
	return nil

}
