package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetCat(c *gin.Context) {
	id := c.Param("id")

	var cat models.Cat

	if result := h.DB.First(&cat, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &cat)
}
