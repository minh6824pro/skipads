package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"github.com/gin-gonic/gin"
)

func (g *GinHttp) HandleGetUserSkipAds(c *gin.Context) {
	// check user id in param
	userID := c.Param("user_id")

	// get user skip ads by user id
	userSkipAds, err := g.query.GetUserSkipAds(c.Request.Context(), userID)
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	resp := httpmodel.Response{}
	resp.Data = userSkipAds
	resp.Message = "get user skip ads successfully"
	c.JSON(200, resp)
}
