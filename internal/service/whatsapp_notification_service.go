package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsAppNotificationService struct {
	store     *sqlstore.Container
	client    *whatsmeow.Client
	clientLog *waLog.Logger
}

func NewWhatsAppNotificationService(
	store *sqlstore.Container,
	client *whatsmeow.Client,
	clientLog *waLog.Logger,
) *WhatsAppNotificationService {

	return &WhatsAppNotificationService{
		store:     store,
		client:    client,
		clientLog: clientLog,
	}
}

func (s *WhatsAppNotificationService) Auth() error {

	if s.client.Store.ID == nil {
		if err := s.handleNewLogin(); err != nil {
			return fmt.Errorf("login failed: %w", err)
		}
	}

	return nil
}

func (s *WhatsAppNotificationService) SendToWhatsApp(companies entity.CompanyCollection, message string) error {

	if err := s.handleExistingSession(companies, message); err != nil {
		return fmt.Errorf("failed to send messages: %w", err)
	}

	s.waitForInterrupt()
	return nil
}

func (s *WhatsAppNotificationService) handleNewLogin() error {
	qrChan, _ := s.client.GetQRChannel(context.Background())
	if err := s.client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	for evt := range qrChan {
		if evt.Event == "code" {
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
		} else {
			fmt.Println("Login event:", evt.Event)
		}
	}

	return nil
}

func (s *WhatsAppNotificationService) handleExistingSession(companies entity.CompanyCollection, message string) error {

	if !s.client.IsConnected() {
		if err := s.client.Connect(); err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}
	}

	for _, company := range companies.Get() {

		// Capture company in loop scope for goroutine
		company := company

		go func() {
			if err := s.sendText(context.Background(), s.client, message, company.PhoneNumber()); err != nil {
				fmt.Printf("Send failed to %s: %v\n", company.Name, err)
			} else {
				fmt.Printf("Message sent successfully to %s!\n", company.Name)
			}
		}()

	}

	return nil
}

func (s *WhatsAppNotificationService) sendText(ctx context.Context, client *whatsmeow.Client, text string, phoneE164Digits string) error {
	jid := types.NewJID(phoneE164Digits, types.DefaultUserServer)
	msg := &waProto.Message{Conversation: &text}
	_, err := client.SendMessage(ctx, jid, msg)
	return err
}

func (s *WhatsAppNotificationService) waitForInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	s.client.Disconnect()
}
