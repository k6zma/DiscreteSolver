package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k6zma/DiscreteSolver/internal/api/middlewares"
	"github.com/k6zma/DiscreteSolver/pkg/api/routers"
)

func main() {
	router := gin.New()

	router.SetTrustedProxies(nil)

	router.Use(middlewares.LoggerMiddleware())

	routers.InitializeRoutes(router)

	router.Run(":8080")
}
