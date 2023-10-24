package main

import (
	// "fmt"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}


func main() {
	r := gin.Default()
	r.POST("/signUp", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth ,controllers.Validate)


	r.Run() // listen and serve on 0.0.0.0:8080
}