package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inabe49/gin-example/controllers"
)

func main() {
	router := gin.New()

	router.GET("/me", controllers.Me)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":9000")
}
