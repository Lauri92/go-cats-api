package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetCat(c *gin.Context) {
	id := c.Param("id")

	var cat models.Cat

	result := h.DB.First(&cat, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &cat)
}
