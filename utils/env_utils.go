package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv загружает переменные окружения из файла .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл, будут использоваться системные переменные окружения")
	}
}

// GetEnv возвращает значение переменной окружения или значение по умолчанию
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
