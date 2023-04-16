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
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var cat models.Cat

	cat.Name = body.Name
	cat.Age = body.Age
	cat.Description = body.Description

	if result := h.DB.Create(&cat); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &cat)
}
