package httpmodel

import (
	"SkipAdsV2/errorcode"
	"errors"
)

type UseSkipAdsRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	AppID       string `json:"app_id" binding:"required"`
	Quantity    uint32 `json:"quantity" binding:"required"`
	Description string `json:"description"`
}

func (req *UseSkipAdsRequest) Validate() error {
	if req.UserID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("user_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.Quantity == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("quantity can't be 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.AppID == "" {
		return &errorcode.ErrorService{
			InternalError: errors.New("app_id can't be nil"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	return nil
}
