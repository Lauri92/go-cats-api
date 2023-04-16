package cats

import (
	"net/http"

	"github.com/Lauri92/go-cats-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetCats(c *gin.Context) {
	var cats []models.Cat

	result := h.DB.Find(&cats)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &cats)
}
