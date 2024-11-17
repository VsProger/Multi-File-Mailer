package main

import (
	"github.com/vsproger/Doodocs-days-2.0/application"
	"github.com/vsproger/Doodocs-days-2.0/config"
	"github.com/vsproger/Doodocs-days-2.0/logger"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	newLogger := logger.NewLogger()

	app := application.NewApplication(cfg, newLogger)

	app.Start()
}
