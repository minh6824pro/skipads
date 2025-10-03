package httpmodel

import (
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"errors"
	"time"
)

type GrantSkipAdsRequest struct {
	UserID        string `json:"user_id" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	Quantity      uint32 `json:"quantity" binding:"required"`
	ExpiresAfter  uint32 `json:"expires_after" binding:"required"`
	Description   string `json:"description"`
}

func (req *GrantSkipAdsRequest) Validate() error {
	if req.UserID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("user_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.Quantity == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("quantity can't be less than 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.TransactionID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("transaction_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.ExpiresAfter <= 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("expires_after can't be less than 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	return nil
}

func (req *GrantSkipAdsRequest) ConvertToEventAddSkipAdsGrant() entities.EventAddSkipAds {

	if req.Description == "" {
		req.Description = "Grant user skip ads"
	}

	now := time.Now()
	evAdd := entities.EventAddSkipAds{
		UserID:        req.UserID,
		SourceEventID: req.TransactionID,
		Description:   req.Description,
		Quantity:      req.Quantity,
		ExpiresAt:     now.Add(time.Duration(req.ExpiresAfter) * 24 * time.Hour),
	}
	return evAdd
}
