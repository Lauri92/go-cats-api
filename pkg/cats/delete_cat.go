package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteCat(c *gin.Context) {
	id := c.Param("id")

	var cat models.Cat

	result := h.DB.First(&cat, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&cat)

	c.Status(http.StatusOK)
}
