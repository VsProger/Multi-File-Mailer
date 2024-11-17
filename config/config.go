package config

import (
	"github.com/vsproger/Doodocs-days-2.0/utils"
)

type Config struct {
	Port         string
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func LoadConfig() (*Config, error) {
	utils.LoadEnv()
	return &Config{
		Port:         utils.GetEnv("PORT", ":8080"),
		SMTPHost:     utils.GetEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     utils.GetEnv("SMTP_PORT", "587"),
		SMTPUsername: utils.GetEnv("SMTP_USERNAME", ""),
		SMTPPassword: utils.GetEnv("SMTP_PASSWORD", ""),
	}, nil
}
