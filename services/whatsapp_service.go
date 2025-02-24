package services

import (
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
	return &WAService{Container: c}
}

func (s *WAService) GetFirstDevice() error {
	device, err := s.Container.GetFirstDevice()
	if err != nil {
		return err
	}
	s.Device = device
	return err
}

func (s *WAService) CreateClient() {
	client := whatsmeow.NewClient(s.Device, nil)
	s.Client = client
}

// func (s *WAService) RegisterUser(name, email string) (*models.User, error) {
// 	user := &models.User{Name: name, Email: email}
// 	err := s.repo.CreateUser(user)
// 	return user, err
// }
