package main

import (
	"gin-api/config"
	"gin-api/db"
	"gin-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	r := gin.Default()
	// Create API version groups
	v1 := r.Group("/api/v1")
	// Register routes
	routes.Routes(v1)

	// userRepo := repositories.NewUserRepository(db.DB)
	// userService := services.NewUserService(userRepo)
	// userHandler := handlers.NewUserHandler(userService)

	// r := routes.SetupRouter(userHandler)
	log.Println("Server running on port", config.PORT)
	r.Run(":" + config.PORT)
}
