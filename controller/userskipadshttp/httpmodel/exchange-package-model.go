package httpmodel

import (
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"errors"
)

type ExchangeRequest struct {
	UserID        string `json:"user_id" binding:"required"`
	PackageID     string `json:"package_id" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	Description   string `json:"description"`
}

func (req *ExchangeRequest) Validate() error {
	if req.UserID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("user_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.PackageID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("package_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.TransactionID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("transaction_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	return nil
}

func (req *ExchangeRequest) ConvertToEventAddSkipAdsExchange() entities.EventAddSkipAds {

	if req.Description == "" {
		req.Description = "User exchange skip ads"
	}
	evAdd := entities.EventAddSkipAds{
		UserID:        req.UserID,
		PackageID:     &req.PackageID,
		SourceEventID: req.TransactionID,
		Description:   req.Description,
	}
	return evAdd
}
