package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	// Passing json:"param" tags to the struct fields
	// allows us to bind JSON data from the request body to these fields.
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	// *context -> Get all the context of the request
	router.GET("/user/:uid", func(context *gin.Context) {
		uid := context.Param("uid")
		context.String(
			http.StatusOK,
			"Hello %s", uid,
		)
	})

	router.GET("/hello", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname") // Shortcut for context.Request.URL.Query().Get("lastname")

		context.String(
			http.StatusOK,
			"Hello %s %s", firstname, lastname,
		)
	})

	router.POST("/welcome", func(context *gin.Context) {

		var user User

		// Check if there is an error while binding the JSON body to the user struct
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{ // If there is an error, return a HTTP Status Code with a message
					http.StatusText(http.StatusBadRequest): "Cannot handle JSON struct correctly!",
				})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Welcome " + user.Name})

	})
	router.Run(":8088")
}
