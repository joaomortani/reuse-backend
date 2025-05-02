package v1

import (
	items "reuse-api/handlers/item"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine, handler *items.ItemHandler) {
	item := r.Group("/api/v1")
	// item.Use(middlewares.AuthMiddleware())
	{
		item.POST("/items", handler.CreateItem)
		item.GET("/items", handler.GetAllItems)
		item.GET("/items/:id", handler.GetItemByID)
	}
}
