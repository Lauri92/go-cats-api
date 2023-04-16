package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddCatRequestBody struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

func (h handler) AddCat(c *gin.Context) {
	body := AddCatRequestBody{}

	// getting request's body
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var cat models.Cat

	cat.Name = body.Name
	cat.Age = body.Age
	cat.Description = body.Description

	result := h.DB.Create(&cat)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &cat)
}
