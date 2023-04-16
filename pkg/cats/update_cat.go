package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateCatRequestBody struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

func (h handler) UpdateCat(c *gin.Context) {
	id := c.Param("id")
	body := UpdateCatRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var cat models.Cat

	if result := h.DB.First(&cat, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	cat.Name = body.Name
	cat.Age = body.Age
	cat.Description = body.Description

	h.DB.Save(&cat)

	c.JSON(http.StatusOK, &cat)
}
