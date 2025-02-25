package utils

import (
	"log"

	"go.mau.fi/whatsmeow"
)

func MustInit[T any](initializer func() (T, error), name string) T {
	instance, err := initializer()
	if err != nil {
		log.Fatalf("Error initializing %s: %v", name, err)
	}
	return instance
}

// Create Newsletter messages params
func CreateNLParams(count, before int) *whatsmeow.GetNewsletterMessagesParams {
	return &whatsmeow.GetNewsletterMessagesParams{
		Count:  count,
		Before: before,
	}
}
