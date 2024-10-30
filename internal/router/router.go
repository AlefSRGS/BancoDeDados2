package router

import (
	"github.com/gin-gonic/gin"
)

func InicializeRouter() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello-World",
		})
	})
	router.Run(":8080")
}
