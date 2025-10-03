package userskipadshttp

import (
	"SkipAdsV2/config"
	"SkipAdsV2/errorcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ValidateHeaderAuthInternal: check internal API key
func (g *GinHttp) ValidateHeaderAuthInternal() gin.HandlerFunc {
	return func(c *gin.Context) {
		internalKey := c.GetHeader(config.HeaderInternalAuth)
		if internalKey == "" || internalKey != g.cfg.GetInternalAPIKey() {
			g.ErrorHandlerCentralized(c, &errorcode.ErrorService{
				InternalError: fmt.Errorf("this request is not authenticated"),
				ErrorCode:     errorcode.ErrAuth})
			return
		}
		c.Next()
	}
}

// AddRequestIDToContext: save X-Request-ID to context
func (g *GinHttp) AddRequestIDToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		c.Set(config.KeyRequestID, requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
