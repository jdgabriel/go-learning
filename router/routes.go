package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/jdgabriel/go-learning/docs"
	"github.com/jdgabriel/go-learning/handler"
)

func InitializeRoutes(router *gin.Engine) {
	// Init handlers
	handler.InitHandlers()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Group routes
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.GetOpeningHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)

	}

	// Init swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
