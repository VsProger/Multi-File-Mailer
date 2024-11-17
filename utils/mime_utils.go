package utils

import (
	"mime/multipart"
	"path/filepath"
)

var allowedArchiveMimeTypes = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/xml": true,
	"image/jpeg":      true,
	"image/png":       true,
}

var allowedEmailMimeTypes = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/pdf": true,
}

func IsValidArchiveMimeType(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	return allowedArchiveMimeTypes[mimeType]
}

func IsValidEmailMimeType(fileHeader *multipart.FileHeader) bool {
	mimeType := fileHeader.Header.Get("Content-Type")
	return allowedEmailMimeTypes[mimeType]
}

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
		return "application/octet-stream"
	}
}
