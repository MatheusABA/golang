package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()

	request.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "world",
		})
	})

	request.Run(":8088")
}
