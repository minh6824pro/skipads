package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (g *GinHttp) HandleExchangePackage(c *gin.Context) {

	// read and validate data from request
	var req httpmodel.ExchangeRequest
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

	evExchange := req.ConvertToEventAddSkipAdsExchange()

	// handle request
	err := g.command.HandleEventExchangePackage(c.Request.Context(), &evExchange)
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	res := httpmodel.Response{
		Message: "Exchange created successfully",
		Data:    evExchange,
	}
	c.JSON(http.StatusOK, res)
}
