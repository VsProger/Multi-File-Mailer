package config

import (
	"github.com/vsproger/Doodocs-days-2.0/utils"
)

type Config struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func LoadConfig() (*Config, error) {
	utils.LoadEnv()
	return &Config{
		SMTPHost:     utils.GetEnv("SMTP_HOST", ""),
		SMTPPort:     utils.GetEnv("SMTP_PORT", ""),
		SMTPUsername: utils.GetEnv("SMTP_USERNAME", ""),
		SMTPPassword: utils.GetEnv("SMTP_PASSWORD", ""),
	}, nil
}
