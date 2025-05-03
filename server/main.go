package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/handlers"
	"main.go/middlewares"
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
	r.POST("/login", controllers.LoginUser)
	r.GET("/profile", middlewares.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Jai Shri RadhaVallabh",
		})
	})

	r.Run(":8000")
}
