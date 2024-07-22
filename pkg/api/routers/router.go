package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/k6zma/DiscreteSolver/internal/api/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/relation-properties", handlers.GetRelationPropertiesHandler)
		api.POST("/generate-relation-graph", handlers.GenerateRelationGraphHandler)
		api.POST("/generate-truth-table", handlers.GenerateTruthTableHandler)
		api.POST("/fixed-length-encode", handlers.FixedLengthEncodeHandler)
		api.POST("/fixed-length-decode", handlers.FixedLengthDecodeHandler)
		api.POST("/shennon-fano-encode", handlers.ShennonFanoEncodeHandler)
		api.POST("/shennon-fano-decode", handlers.ShennonFanoDecodeHandler)
		api.POST("/create-venn-diagram", handlers.CreateVennDiagramHandler)
	}
}
