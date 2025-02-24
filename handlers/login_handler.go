package handlers

import (
	"gin-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service *services.WAService
}

func NewLoginHandler(service *services.WAService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (h *LoginHandler) LoginPairPhone(c *gin.Context) {
	phoneNumber := c.Query("phone")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone number is required"})
		return
	}
	log.Println("Received pairing request for:", phoneNumber)
	h.service.Client.PairPhone(phoneNumber, true, 1, "Chrome (mac)")
}

// func (h *LoginHandler) CreateUser(c *gin.Context) {
// 	var req models.User
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := h.service.RegisterUser(req.Name, req.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }
