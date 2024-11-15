package utils

import (
	"mime/multipart"
	"path/filepath"
)

// Allowed MIME-типы для архивирования
var allowedArchiveMimeTypes = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/xml": true,
	"image/jpeg":      true,
	"image/png":       true,
}

// Allowed MIME-типы для отправки по почте
var allowedEmailMimeTypes = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/pdf": true,
}

// IsValidArchiveMimeType проверяет MIME-тип файла для архивирования
func IsValidArchiveMimeType(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	return allowedArchiveMimeTypes[mimeType]
}

// IsValidEmailMimeType проверяет MIME-тип файла для отправки по почте
func IsValidEmailMimeType(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	return allowedEmailMimeTypes[mimeType]
}

// DetectMimeType определяет MIME-тип файла на основе его расширения, если заголовок отсутствует
func DetectMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".xml":
		return "application/xml"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream" // Стандартный тип для неизвестных файлов
	}
}
