package userskipadshttp

import (
	"SkipAdsV2/controller/userskipadshttp/httpmodel"
	"SkipAdsV2/errorcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (g *GinHttp) HandleGetUserSkipAds(c *gin.Context) {
	// check user id in param
	idStr := c.Param("user_id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		g.ErrorHandlerCentralized(c, &errorcode.ErrorService{
			InternalError: err,
			ErrorCode:     errorcode.ErrInvalidRequest,
		})
		return
	}

	// get user skip ads by user id
	userSkipAds, err := g.query.GetUserSkipAds(c.Request.Context(), int32(userID))
	if err != nil {
		g.ErrorHandlerCentralized(c, err)
		return
	}
	resp := httpmodel.Response{}
	resp.Data = userSkipAds
	resp.Message = "get user skip ads successfully"
	c.JSON(200, resp)
}
