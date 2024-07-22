package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k6zma/DiscreteSolver/internal/api/models"
	"github.com/k6zma/DiscreteSolver/internal/mathalgos"
)

func GenerateTruthTableHandler(c *gin.Context) {
	var request models.GenerateTruthTableRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	generator := mathalgos.NewTruthTableGenerator(request.Expression)
	imageData, err := generator.CreateTruthTableImage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", imageData)
}
