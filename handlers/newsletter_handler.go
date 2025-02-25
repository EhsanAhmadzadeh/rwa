package handlers

import "gin-api/services"

type NewsletterHandler struct {
	service *services.WAService
}

func NewNewsletterHandler(service *services.WAService) *NewsletterHandler {
	return &NewsletterHandler{service}
}
