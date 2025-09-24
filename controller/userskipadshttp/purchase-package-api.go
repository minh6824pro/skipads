package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (g *GinHttp) HandlePurchasePackage(c *gin.Context) {

	// read and validate data from request
	var req httpmodel.PurchaseRequest
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

	evPurchase := req.ConvertToEventAddSkipAdsPurchase()

	// handle request
	err := g.command.HandleEventPurchasePackage(c.Request.Context(), &evPurchase)
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	res := httpmodel.Response{
		Message: "Purchase created successfully",
		Data:    evPurchase,
	}
	c.JSON(http.StatusOK, res)
}
