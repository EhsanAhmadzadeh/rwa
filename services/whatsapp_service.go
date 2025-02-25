package services

import (
	"fmt"
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

type WAService struct {
	Container *sqlstore.Container
	Device    *store.Device
	Client    *whatsmeow.Client
}

func NewWAService(c *sqlstore.Container) *WAService {
	s := &WAService{Container: c}

	// Ensure the device is initialized before creating the client
	if err := s.GetFirstDevice(); err != nil {
		log.Println("Failed to get first device:", err)
		return s
	}
	s.GetFirstDevice()
	s.CreateClient()
	return s
}

func (s *WAService) GetFirstDevice() error {
	device, err := s.Container.GetFirstDevice()
	if err != nil {
		log.Println("Error retrieving first device:", err)
		return err
	}

	if device == nil {
		log.Println("No device found in the database")
		return fmt.Errorf("no device found")
	}

	s.Device = device
	return nil
}

func (s *WAService) CreateClient() {
	if s.Device == nil {
		log.Println("Error: Device is nil, cannot create WhatsApp client")
		return
	}
	// Connect to the WhatsApp WebSocket

	client := whatsmeow.NewClient(s.Device, nil)
	s.Client = client

	// Only connect  WhatsApp WebSocket if not already connected
	if !s.Client.IsConnected() {
		if err := s.Client.Connect(); err != nil {
			fmt.Println("Failed to connect to WhatsApp: %w", err)
		}
	}

	log.Println("WhatsApp client successfully connected to WebSocket")
}

// func (s *WAService) RegisterUser(name, email string) (*models.User, error) {
// 	user := &models.User{Name: name, Email: email}
// 	err := s.repo.CreateUser(user)
// 	return user, err
// }
