package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k6zma/DiscreteSolver/internal/api/models"
	"github.com/k6zma/DiscreteSolver/internal/mathalgos"
)

func GetRelationPropertiesHandler(c *gin.Context) {
	var model models.BinaryRelationModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	properties := mathalgos.GetRelationProperties(model)
	c.JSON(http.StatusOK, gin.H{"properties": properties})
}

func GenerateRelationGraphHandler(c *gin.Context) {
	var model models.BinaryRelationModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	graph := mathalgos.NewBinaryRelationGraph(model)
	imageData, err := graph.GenerateImage()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "image/png", imageData)
}
