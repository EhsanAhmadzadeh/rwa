package handlers

import (
	"fmt"
	"gin-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service *services.WAService
}

func NewLoginHandler(service *services.WAService) *LoginHandler {
	return &LoginHandler{service}
}

func (h *LoginHandler) LoginPairPhone(c *gin.Context) {

	if h.service.Client == nil {
		log.Println("Error: WhatsApp Client is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WhatsApp client is not initialized"})
		return
	}

	phoneNumber := c.Query("phone")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone number is required"})
		return
	}
	log.Println("Received pairing request for:", phoneNumber)
	pairingCode, err := h.service.Client.PairPhone(phoneNumber, true, 1, "Chrome (mac)")
	if err != nil {
		log.Println("Pairing failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Pairing code sent:", pairingCode)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Please enter the pairing code %s for account with phone number %s", pairingCode, phoneNumber),
	})
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
