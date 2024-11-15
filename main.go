package main

import (
	"github.com/vsproger/Doodocs-days-2.0/application"
	"github.com/vsproger/Doodocs-days-2.0/config"
	"github.com/vsproger/Doodocs-days-2.0/logger"
	"log"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация логгера
	newLogger := logger.NewLogger()

	// Инициализация приложения
	app := application.NewApplication(cfg, newLogger)

	// Настройка маршрутов
	app.SetupRoutes()

	// Запуск сервера
	app.Start()
}
