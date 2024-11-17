package services

import (
	"bytes"
	"fmt"
	"github.com/vsproger/Doodocs-days-2.0/config"
	"io"
	"mime/multipart"
	"net/smtp"
	"strings"
)

func SendEmailWithAttachment(to []string, file multipart.File, filename, mimeType string, cfg *config.Config) error {
	if cfg.SMTPUsername == "" {
		return fmt.Errorf("пустое имя пользователя")
	}
	from := cfg.SMTPUsername
	subject := "Sending File Attachment"
	body := "Please find the attached file."
	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: multipart/mixed; boundary=boundary\n\n" +
		"--boundary\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\n\n" +
		body + "\n\n" +
		"--boundary\n" +
		"Content-Type: " + mimeType + "\n" +
		"Content-Disposition: attachment; filename=\"" + filename + "\"\n\n"

	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	msgBuffer := bytes.NewBufferString(msg)
	msgBuffer.Write(fileData)
	msgBuffer.WriteString("\n--boundary--")

	auth := smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPHost)
	err = smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, from, to, msgBuffer.Bytes())
	if err != nil {
		return fmt.Errorf("ошибка отправки письма: %v", err)
	}

	return nil
}
