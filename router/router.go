package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	router := gin.Default()
	InitializeRoutes(router)
	// listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
