package service

import (
	"context"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"github.com/victoratsuta/google_map2whatsapp/internal/entity"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var wg = &sync.WaitGroup{}

type messageSendReport struct {
	company string
	err     error
}

type WhatsAppNotificationService struct {
	store     *sqlstore.Container
	client    *whatsmeow.Client
	clientLog waLog.Logger
}

func NewWhatsAppNotificationService(
	store *sqlstore.Container,
	client *whatsmeow.Client,
	clientLog waLog.Logger,
) *WhatsAppNotificationService {
	return &WhatsAppNotificationService{
		store:     store,
		client:    client,
		clientLog: clientLog,
	}
}

func (s *WhatsAppNotificationService) Auth() error {
	if s.client.Store.ID == nil {
		if err := s.handleNewQRScanLogin(); err != nil {
			return fmt.Errorf("login to Whatsapp is failed: %w", err)
		}
	}

	return nil
}

func (s *WhatsAppNotificationService) SendToWhatsApp(companies entity.CompanyCollection, message string) error {
	results := make(chan messageSendReport, len(companies.Get()))
	if !s.client.IsConnected() {
		if err := s.client.Connect(); err != nil {
			return fmt.Errorf("failed to connect: %w", err)
		}
	}

	defer s.client.Disconnect()

	for _, company := range companies.Get() {
		wg.Add(1)
		go func() {
			if err := s.sendText(s.client, message, company.PhoneNumber()); err != nil {
				results <- messageSendReport{
					company: company.Name(),
					err:     fmt.Errorf("send failed to %s: %w", company.Name(), err),
				}
			} else {
				results <- messageSendReport{
					company: company.Name(),
					err:     nil,
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(results)
	var errs []error
	for r := range results {
		if r.err != nil {
			errs = append(errs, fmt.Errorf("send to %s: %w", r.company, r.err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("error report during send to whatsapp: %v", errs)
	}

	return nil
}

func (s *WhatsAppNotificationService) handleNewQRScanLogin() error {
	qrChan, _ := s.client.GetQRChannel(context.TODO())
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

func (*WhatsAppNotificationService) sendText(client *whatsmeow.Client, text, phoneE164Digits string) error {
	_, err := client.SendMessage(
		context.TODO(),
		types.NewJID(phoneE164Digits, types.DefaultUserServer),
		&waE2E.Message{Conversation: &text},
	)
	return err
}
