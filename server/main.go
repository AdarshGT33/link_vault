package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/handlers"
)

func init() {
	handlers.LoadEnv()
	handlers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shri hari",
		})
	})
	r.POST("/signup", controllers.RegisterUser)

	r.Run(":8000")
}
