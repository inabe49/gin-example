package controllers

import "github.com/gin-gonic/gin"

func Me(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "user",
	})
}
