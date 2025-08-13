package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "world",
		})
	})

	router.Run(":8088")
}
