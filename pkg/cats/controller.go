package cats

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/cats")
	routes.POST("/", h.AddCat)
	routes.GET("/", h.GetCats)
	routes.GET("/:id", h.GetCat)
	routes.PUT("/:id", h.UpdateCat)
	routes.DELETE("/:id", h.DeleteCat)
}
