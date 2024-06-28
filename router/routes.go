package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jdgabriel/go-learning/handler"
)

func InitializeRoutes(router *gin.Engine) {
	// Init handlers
	handler.InitHandlers()

	// Group routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}
}
