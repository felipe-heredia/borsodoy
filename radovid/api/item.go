package api

import (
	"borsodoy/radovid/internal/models"
	"borsodoy/radovid/internal/service"
	"borsodoy/radovid/pkg/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var newItemData models.CreateItem

	if err := c.BindJSON(&newItemData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := service.CreateItem(newItemData)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, newItem)
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		panic("error")
	}
	item, err := service.GetItemById(id)

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, item)
}

func GetItems(c *gin.Context) {
  items, err := service.GetItems()

	if err != nil {
		if httpError, ok := err.(*utility.HttpError); ok {
			c.JSON(httpError.Status, gin.H{"message": httpError.Message})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, items)
}
