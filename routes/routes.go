package routes

import (
	"fmt"
	"gin-api/db"
	"gin-api/handlers"
	"gin-api/services"
	"gin-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	Container := db.DB
	service := services.NewWAService(Container)
	LoginHandler := handlers.NewLoginHandler(service)
	router.POST("/login", LoginHandler.LoginPairPhone)
	router.POST("/get-nl-messages", func(ctx *gin.Context) {
		link := ctx.Query("link")
		metadata, err := service.Client.GetNewsletterInfoWithInvite(link)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to get newsletter info.",
				"error":   fmt.Sprint(err),
			})
		} else {
			c, _ := strconv.Atoi(ctx.Query("count"))
			b, _ := strconv.Atoi(ctx.Query("before"))
			param := utils.CreateNLParams(c, b)
			msgs, _ := service.Client.GetNewsletterMessages(metadata.ID, param)
			ctx.JSON(http.StatusOK, gin.H{
				"metadata": msgs,
			})

		}
	})
}
