package config

import (
	"context"
	"fmt"

	"github.com/victoratsuta/google_map2whatsapp/internal/repo"
	"github.com/victoratsuta/google_map2whatsapp/internal/service"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type Container struct {
	config *Config

	companiesRepo repo.CompaniesRepositoryInterface

	whatsappService service.WhatsAppNotificationServiceInterface
}

func NewContainer(config *Config) (*Container, error) {
	c := &Container{
		config: config,
	}

	if err := c.initRepositories(); err != nil {
		return nil, fmt.Errorf("failed to initialize repositories: %w", err)
	}

	if err := c.initServices(); err != nil {
		return nil, fmt.Errorf("failed to initialize services: %w", err)
	}

	return c, nil
}

func (c *Container) initRepositories() error {

	if c.config.App.Env == "prod" {
		c.companiesRepo = repo.NewGoogleMapsCompaniesRepository(c.config.GoogleMap.ApiKey)
	} else {
		c.companiesRepo = repo.NewCompaniesRepositoryStub(c.config.GoogleMap.ApiKey)
	}

	return nil
}

func (c *Container) initServices() error {

	dbLog := waLog.Stdout("Database", c.config.Log.Level, true)
	clientLog := waLog.Stdout("Client", c.config.Log.Level, true)
	ctx := context.Background()
	store, err := sqlstore.New(ctx, "sqlite3", "file:cache/examplestore.db?_foreign_keys=on", dbLog)

	if err != nil {
		panic(err)
	}

	deviceStore, _ := store.GetFirstDevice(ctx)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	c.whatsappService = service.NewWhatsAppNotificationService(
		store,
		client,
		&ctx,
		&clientLog,
	)
	return nil
}

func (c *Container) GetCompaniesRepository() repo.CompaniesRepositoryInterface {
	return c.companiesRepo
}

func (c *Container) GetWhatsAppService() service.WhatsAppNotificationServiceInterface {
	return c.whatsappService
}
