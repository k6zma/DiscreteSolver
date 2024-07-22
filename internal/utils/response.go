package utils

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
