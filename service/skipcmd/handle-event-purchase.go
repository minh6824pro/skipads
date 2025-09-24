package skipcmd

import (
	"SkipAdsV2/entities"
	"context"
	"time"
)

func (cmd *Command) HandleEventPurchasePackage(ctx context.Context, eventPurchase *entities.EventAddSkipAds) error {
	// TODO check user exists

	// get package info
	pkg, err := cmd.db.GetPurchasePackageByID(ctx, eventPurchase.PackageID)
	if err != nil {
		return err
	}
	eventPurchase.Quantity = pkg.Quantity
	eventPurchase.QuantityUsed = 0
	eventPurchase.Type = entities.EventAddSkipAdsPurchase
	expires := time.Now().Add(time.Duration(pkg.ExpiresAfter) * 24 * time.Hour)
	eventPurchase.ExpiresAt = expires

	// create event
	err = cmd.db.CreateEventAddSkipAds(ctx, eventPurchase)
	if err != nil {
		return err
	}
	return nil

}
