package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (g *GinHttp) HandleGrantSkipAds(c *gin.Context) {

	// read and validate data from request
	var req httpmodel.GrantSkipAdsRequest
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

	evGrant := req.ConvertToEventAddSkipAdsGrant()

	// handle request
	err := g.command.HandleEventGrantSkipAds(c.Request.Context(), &evGrant)
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	res := httpmodel.Response{
		Message: "Grant skip ads successfully",
		Data:    evGrant,
	}
	c.JSON(http.StatusOK, res)
}
