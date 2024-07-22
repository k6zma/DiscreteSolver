package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/k6zma/DiscreteSolver/internal/api/models"
	"github.com/k6zma/DiscreteSolver/internal/mathalgos"
)

func init() {
	godotenv.Load()
}

func CreateVennDiagramHandler(c *gin.Context) {
	var request models.VennDiagramRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiKey := os.Getenv("WOLFRAM_ALPHA_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is missing"})
		return
	}

	vennDiagramBuilder := mathalgos.NewVennDiagramBuilder(apiKey)
	imageURL, err := vennDiagramBuilder.BuildDiagram(request.Expression)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.VennDiagramResponse{
		ImageURL: imageURL,
	}
	c.JSON(http.StatusOK, response)
}
