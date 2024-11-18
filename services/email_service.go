package services

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/vsproger/Doodocs-days-2.0/config"
	"io"
	"mime/multipart"
	"net/smtp"
)

func SendEmailWithAttachment(to []string, file multipart.File, filename, mimeType string, cfg *config.Config) error {
	if cfg.SMTPUsername == "" {
		return fmt.Errorf("пустое имя пользователя")
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	e := email.NewEmail()
	e.From = cfg.SMTPUsername
	e.To = to
	e.Subject = "Sending File Attachment"
	e.Text = []byte("Please find the attached file.")

	e.Attach(bytes.NewReader(fileData), filename, mimeType)

	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	auth := smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPHost)

	err = e.Send(addr, auth)
	if err != nil {
		return fmt.Errorf("ошибка отправки письма: %v", err)
	}

	return nil
}
