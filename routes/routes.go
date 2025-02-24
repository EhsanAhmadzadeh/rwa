package routes

import (
	"gin-api/db"
	"gin-api/handlers"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	Container := db.DB
	service := services.NewWAService(Container)
	LoginHandler := handlers.NewLoginHandler(service)
	router.POST("/login", LoginHandler.LoginPairPhone)
}
