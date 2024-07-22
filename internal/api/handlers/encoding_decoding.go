package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k6zma/DiscreteSolver/internal/api/models"
	"github.com/k6zma/DiscreteSolver/internal/mathalgos"
)

func FixedLengthEncodeHandler(c *gin.Context) {
	var request models.EncodeDecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encoder := mathalgos.NewFixedLengthCoding(request.String)
	encodedString := encoder.Encode(request.String)

	response := models.EncodeResponse{
		EncodedString:     encodedString,
		Alphabet:          encoder.GetAlphabetDict(),
		AverageCodeLength: float64(encoder.AverageCodeLength()),
	}
	c.JSON(http.StatusOK, response)
}

func FixedLengthDecodeHandler(c *gin.Context) {
	var request models.DecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decoder := mathalgos.RecreateFromAlphabet(request.Alphabet)
	decodedString := decoder.Decode(request.EncodedString)

	response := models.DecodeResponse{
		DecodedString: decodedString,
	}
	c.JSON(http.StatusOK, response)
}

func ShennonFanoEncodeHandler(c *gin.Context) {
	var request models.EncodeDecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encoder := mathalgos.NewShennonFanoCoding(request.String)
	encodedString := encoder.Encode(request.String)

	response := models.EncodeResponse{
		EncodedString:     encodedString,
		Alphabet:          encoder.GetAlphabetDict(),
		AverageCodeLength: encoder.AverageCodeLength(),
	}
	c.JSON(http.StatusOK, response)
}

func ShennonFanoDecodeHandler(c *gin.Context) {
	var request models.DecodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decoder := mathalgos.RecreateFromCodes(request.Alphabet)
	decodedString := decoder.Decode(request.EncodedString)

	response := models.DecodeResponse{
		DecodedString: decodedString,
	}
	c.JSON(http.StatusOK, response)
}
