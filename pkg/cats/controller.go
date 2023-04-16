package cats

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterCatRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	catRoutes := r.Group("/cats")
	catRoutes.POST("/", h.AddCat)
	catRoutes.GET("/", h.GetCats)
	catRoutes.GET("/:id", h.GetCat)
	catRoutes.PUT("/:id", h.UpdateCat)
	catRoutes.DELETE("/:id", h.DeleteCat)
}
