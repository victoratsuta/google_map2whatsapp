package service

import "github.com/victoratsuta/google_map2whatsapp/internal/entity"

type WhatsAppNotificationServiceInterface interface {
	Auth() error
	SendToWhatsApp(companies entity.CompanyCollection, message string) error
}
