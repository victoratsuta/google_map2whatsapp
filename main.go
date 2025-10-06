package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/victoratsuta/google_map2whatsapp/cmd"
	"github.com/victoratsuta/google_map2whatsapp/config"
)

func main() {

	cfg := setupConfig()
	container, _ := config.NewContainer(cfg)

	cmd.Execute(container)
}

func setupConfig() *config.Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	return cfg
}
