package httpmodel

import (
	"SkipAdsV2/errorcode"
	"errors"
)

type UseSkipAdsRequest struct {
	UserID      uint32 `json:"user_id" binding:"required"`
	Quantity    uint32 `json:"quantity" binding:"required"`
	Description string `json:"description"`
}

func (req *UseSkipAdsRequest) Validate() error {
	if req.UserID == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("user_id can't be 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	if req.Quantity == 0 {
		return &errorcode.ErrorService{
			InternalError: errors.New("quantity can't be 0"),
			ErrorCode:     errorcode.ErrInvalidRequest,
		}
	}
	return nil
}
