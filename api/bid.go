package api

import (
	"radovid/internal/models"
	"radovid/internal/service"
	"radovid/pkg/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBid(c *gin.Context) {
	var newBidData models.CreateBid

	if err := c.BindJSON(&newBidData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBid, err := service.CreateBid(newBidData)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bid"})
		return
	}

	c.JSON(http.StatusOK, newBid)
}

func WithdrawnBid(c *gin.Context) {
	id := c.Param("id")

	withdrawn, err := service.WithdrawnBid(id)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}
	}

	if withdrawn == true {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	} else {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
}
