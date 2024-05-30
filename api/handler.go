package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldHandler responds with "Hello, World!"
func HelloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
