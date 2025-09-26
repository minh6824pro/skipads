package skipcmd

import (
	"SkipAdsV2/entities"
	"context"
	"time"
)

func (cmd *Command) HandleEventExchangePackage(ctx context.Context, eventExchange *entities.EventAddSkipAds) error {
	// TODO check user exists

	// get package info
	pkg, err := cmd.db.GetExchangePackageByID(ctx, eventExchange.PackageID)
	if err != nil {
		return err
	}
	eventExchange.Quantity = pkg.Quantity
	eventExchange.QuantityUsed = 0
	eventExchange.Type = entities.EventAddSkipAdsExchange
	expires := time.Now().Add(time.Duration(pkg.ExpiresAfter) * 24 * time.Hour)
	eventExchange.ExpiresAt = expires
	eventExchange.SetPriority()

	// create event
	err = cmd.db.CreateEventAddSkipAds(ctx, eventExchange)
	if err != nil {
		return err
	}
	return nil

}
