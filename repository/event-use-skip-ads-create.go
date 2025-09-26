package repository

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"SkipAdsV2/repository/repomodel"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

func (r *RepoMySQL) ProcessEventUseSkipAds(ctx context.Context, request httpmodel.UseSkipAdsRequest) error {
	userID := request.UserID
	quantity := request.Quantity

	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// Get Usable Skip Ads
	now := time.Now()
	var availableEventAdd []repomodel.AvailableSkipAds
	err := tx.Raw(`
	SELECT id as event_add_id,quantity-quantity_used as remaining
	FROM event_add_skip_ads 
	WHERE user_id= ?
	AND expires_at> ?
 	AND quantity-quantity_used>0
	ORDER BY 
		priority ASC,
		expires_at ASC
	LIMIT 3
	FOR UPDATE`, userID, now).Scan(&availableEventAdd).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	totalAvailableSkipAds := uint32(0)
	for _, event := range availableEventAdd {
		totalAvailableSkipAds += event.Remaining
	}
	if totalAvailableSkipAds < quantity {
		tx.Rollback()
		return &errorcode.ErrorService{
			InternalError: errors.New("insufficient skip ads"),
			ErrorCode:     errorcode.ErrUserSkipAdsInsufficient,
		}
	}

	remainingToUse := quantity
	for _, event := range availableEventAdd {
		if remainingToUse == 0 {
			break
		}
		// calculate usable quantity from each event_add
		useFromThis := min(remainingToUse, event.Remaining)
		if useFromThis > 0 {
			// update event_add.quantity_used
			result := tx.Model(entities.EventAddSkipAds{}).
				Where("id=? AND quantity - quantity_used >= ?", event.EventAddID, useFromThis).
				Update("quantity_used", gorm.Expr("quantity_used + ?", useFromThis))
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}

			if result.RowsAffected == 0 {
				tx.Rollback()
				return &errorcode.ErrorService{
					InternalError: errors.New("concurrent usage detected -  insufficient skip ads"),
					ErrorCode:     errorcode.ErrSystem,
				}
			}

			// create event_sub
			eventSub := entities.EventSubSkipAds{
				UserID:        userID,
				SourceSubID:   event.EventAddID,
				SourceSubType: entities.SourceAddSkipAds,
				QuantityUsed:  useFromThis,
				Type:          entities.EventSubSkipAdsUse,
				Description:   request.Description,
			}

			err = tx.Create(&eventSub).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
		remainingToUse -= useFromThis
	}

	if remainingToUse > 0 {
		tx.Rollback()
		return &errorcode.ErrorService{
			InternalError: errors.New("failed to allocated all use skip ads"),
			ErrorCode:     errorcode.ErrSystem,
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
