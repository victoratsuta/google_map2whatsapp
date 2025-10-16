package config

import (
	"context"
	"fmt"

	"github.com/victoratsuta/google_map2whatsapp/internal/repo"
	"github.com/victoratsuta/google_map2whatsapp/internal/service"
	"github.com/victoratsuta/google_map2whatsapp/pkg/google_maps"
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
		var googleMapsHttpApiClient google_maps.GoogleMapsApiClientInterface
		googleMapsHttpApiClient = google_maps.NewGoogleMapsHttpApiClient(
			c.config.GoogleMap.ApiKey,
			"https://places.googleapis.com/v1/places:searchText",
		)
		c.companiesRepo = repo.NewGoogleMapsCompaniesRepository(googleMapsHttpApiClient)
	} else {
		c.companiesRepo = repo.NewCompaniesRepositoryStub()
	}

	return nil
}

func (c *Container) initServices() error {

	dbLog := waLog.Stdout("Database", c.config.Log.Level, true)
	clientLog := waLog.Stdout("Client", c.config.Log.Level, true)
	store, err := sqlstore.New(context.Background(), "sqlite3", "file:cache/examplestore.db?_foreign_keys=on", dbLog)

	if err != nil {
		panic(err)
	}

	deviceStore, _ := store.GetFirstDevice(context.Background())
	client := whatsmeow.NewClient(deviceStore, clientLog)

	c.whatsappService = service.NewWhatsAppNotificationService(
		store,
		client,
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
