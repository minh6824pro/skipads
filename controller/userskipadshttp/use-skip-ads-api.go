package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (g *GinHttp) HandleUseSkipAds(c *gin.Context) {

	// read and validate data from request
	var req httpmodel.UseSkipAdsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		g.ErrorHandlerCentralized(c, &errorcode.ErrorService{
			InternalError: err,
			ErrorCode:     errorcode.ErrInvalidRequest,
		})
		return
	}
	if err := req.Validate(); err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}

	// handle request
	err := g.command.HandleEventUseSkipAds(c.Request.Context(), req)
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	res := httpmodel.Response{
		Message: "Purchase created successfully",
		Data:    "success",
	}
	c.JSON(http.StatusOK, res)
}
