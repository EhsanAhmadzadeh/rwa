package main

import (
	"fmt"
	"log"
	"net/http"
	"wa-service/config"
	"wa-service/utils"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"

	"github.com/gin-gonic/gin"
)

type WhatsAppClient struct {
	Client *whatsmeow.Client
}

// PairingCode represents the pairing code returned by WhatsApp.
type PairingCode string

// NewContainer initializes and returns a new database container.
// This should only get instantiated in app startup.
// TODO: Use Singletone here!!
func NewContainer() (*sqlstore.Container, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	sqlContainer, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", config.AppConfig.DB_PATH), dbLog)
	if err != nil {
		return nil, err
	}
	config.AppConfig.Container = sqlContainer
	return sqlContainer, nil
}

// Initializes NewWhatsAppClient a new WhatsApp client and set it on the global config.
func NewWhatsAppClient() (*WhatsAppClient, error) {

	deviceStore, err := config.AppConfig.Container.GetFirstDevice()
	if err != nil {
		return nil, fmt.Errorf("failed to get device store: %w", err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	return &WhatsAppClient{Client: client}, nil
}

// PairPhone pairs the WhatsApp client with the given phone number and returns the pairing code.
func (wac *WhatsAppClient) PairPhone(phoneNumber string) (PairingCode, error) {
	// Only connect if not already connected
	if !wac.Client.IsConnected() {
		if err := wac.Client.Connect(); err != nil {
			return "", fmt.Errorf("failed to connect to WhatsApp: %w", err)
		}
	}

	// Pair the phone
	pairingCode, err := wac.Client.PairPhone(phoneNumber, true, 1, "Chrome (mac)")
	if err != nil {
		return "", fmt.Errorf("error sending OTP: %w", err)
	}

	return PairingCode(pairingCode), nil
}

// Close shuts down the WhatsApp client properly.
func (wac *WhatsAppClient) Close() {
	wac.Client.Disconnect()
}

// loginHandler handles the pairing request.
func (wac *WhatsAppClient) loginHandler(c *gin.Context) {
	phoneNumber := c.Query("phone")
	if phoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "phone number is required"})
		return
	}

	log.Println("Received pairing request for:", phoneNumber)

	pairingCode, err := wac.PairPhone(phoneNumber)
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

func main() {
	config.InitConfig()

	config.AppConfig.Container = utils.MustInit(NewContainer, "database container")
	wac := utils.MustInit(NewWhatsAppClient, "WhatsApp client")
	defer wac.Close() // Ensure cleanup on exit
	// config.AppConfig.Container = container

	// Set up Gin router
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "It is working!"})
		})
		v1.POST("/pair-phone", wac.loginHandler)
	}

	// Start the server
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
