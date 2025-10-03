package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (g *GinHttp) ErrorHandlerCentralized(c *gin.Context, err error) {
	// default status request error 500 (internal error)
	status := http.StatusInternalServerError
	var errorCode = errorcode.ErrSystem

	res := httpmodel.Response{}
	// description error for dev
	res.Metadata = err.Error()

	// handle built-in Gin errors
	if ginErr, ok := err.(gin.Error); ok {
		// mapping error code (statusCode) with message define
		if ginErr.Meta != nil {
			if code, ok := ginErr.Meta.(int); ok && code != 0 {
				status = code
			}
		}
		errorCode = errorcode.ErrUnknown
	}
	// handle error in service
	var errSer *errorcode.ErrorService
	if errors.As(err, &errSer) {
		errorCode = errSer.ErrorCode
		// override status by mapping status and error code
		if statusMap := GetStatusByErrCode(errorCode.GetErrCode()); statusMap != nil {
			status = *statusMap
		}
	}

	res.Reason = errorCode.GetErrCode()
	res.Message = errorCode.GetMessage()

	if c.Request.Context().Err() != nil {
		// overwrite status when client closes the connection
		status = 499
	}

	c.JSON(status, res)
	c.Abort()
}
