package main

import (
	"fmt"
	"log"
	"net/http"
	"wa-service/config"
	"wa-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
	err := models.ConnectDB()
	println(err)

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "it is working!"})

		})
		// v1.GET("person/:id", getPersonById)
		// v1.POST("person", addPerson)
		// v1.PUT("person/:id", updatePerson)
		// v1.DELETE("person/:id", deletePerson)
		// v1.OPTIONS("person", options)
	}
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {

		log.Printf("Server running on port %s...", cfg.Port)
	}

}
