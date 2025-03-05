package main

import (
	"go-mysql-api/database"
	"go-mysql-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Connect to Database
	database.ConnectDB()

	// ✅ Initialize Gin Router
	r := gin.Default()

	// ✅ Define Routes
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// ✅ Start Server
	r.Run(":8080")
}
