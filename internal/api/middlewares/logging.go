package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqURI := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		statusColor := getStatusColor(statusCode)

		log.Printf("| \033[%s%3d\033[0m | \033[35;1m%13v\033[0m | \033[36;1m%15s\033[0m | \033[33;1m%s\033[0m | \033[32;1m%s\033[0m |",
			statusColor, statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}

func getStatusColor(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return "32;1m"
	case statusCode >= 300 && statusCode < 400:
		return "33;1m"
	case statusCode >= 400 && statusCode < 500:
		return "31;1m"
	default:
		return "37;1m"
	}
}
