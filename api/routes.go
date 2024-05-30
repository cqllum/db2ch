package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define the hello world route
	router.GET("/hello", HelloWorldHandler)

	return router
}
