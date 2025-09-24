package httpmodel

import (
	"SkipAdsV2/entities"
	"SkipAdsV2/errorcode"
	"errors"
)

type ExchangeRequest struct {
	UserID        uint32 `json:"user_id" binding:"required"`
	PackageID     uint32 `json:"package_id" binding:"required"`
	TransactionID uint32 `json:"transaction_id" binding:"required"`
	Description   string `json:"description"`
}

func (req *ExchangeRequest) Validate() error {
	if req.UserID == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("user_id can't be 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.PackageID == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("package_id can't be 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.TransactionID == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("transaction_id can't be 0"),
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
