package userskipadshttp

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func (g *GinHttp) LoggingRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		req := c.Request

		// continue handle request
		c.Next()

		// after handler finish
		stop := time.Now()
		latency := stop.Sub(start)

		status := c.Writer.Status()
		idRequest := req.Header.Get("X-Request-ID")
		if idRequest == "" {
			idRequest = c.Writer.Header().Get("X-Request-ID")
		}

		byteIn, errParse := strconv.ParseUint(req.Header.Get("Content-Length"), 10, 64)
		if errParse != nil {
			byteIn = 0
		}

		fieldLog := []zap.Field{
			zap.String("request_id", idRequest),
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.String("host", req.Host),
			zap.Int("status", status),
			zap.String("remote_ip", c.ClientIP()),
			zap.String("user_agent", req.UserAgent()),
			zap.Float64("latency", latency.Seconds()),
			zap.String("latency_human", latency.String()),
			zap.Uint64("byte_in", byteIn),
		}

		lg := g.logger
		if status < http.StatusBadRequest {
			lg.Info("api call success", fieldLog...)
		} else if status >= http.StatusBadRequest && status < http.StatusInternalServerError {
			lg.Warn("api call failed by bad request", fieldLog...)
		} else if status >= http.StatusInternalServerError {
			if len(c.Errors) > 0 {
				fieldLog = append(fieldLog, zap.Error(c.Errors.Last().Err))
			}
			lg.Error("api call failed by internal error", fieldLog...)
		}
	}
}
