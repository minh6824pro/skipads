package userskipadshttp

import (
	"SkipAdsV2/errorcode"
	"net/http"
)

var MappingErrorCodeWithStatus = map[string]int{
	errorcode.CodeErrSystem:                  http.StatusInternalServerError,
	errorcode.CodeErrAuth:                    http.StatusUnauthorized,
	errorcode.CodeErrDataRequestInvalid:      http.StatusBadRequest,
	errorcode.CodeErrUserSkipAdsInsufficient: http.StatusBadRequest,
}

func GetStatusByErrCode(errorCode string) *int {
	if status, ok := MappingErrorCodeWithStatus[errorCode]; ok {
		return &status
	}
	return nil
}
