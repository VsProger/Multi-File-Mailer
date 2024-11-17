package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/vsproger/Doodocs-days-2.0/models"
	"github.com/vsproger/Doodocs-days-2.0/utils"
	"io"
	"mime/multipart"
)

func ProcessArchive(file multipart.File, filename string, archiveSize int64) (models.ArchiveInfo, error) {
	zipReader, err := zip.NewReader(file, archiveSize)
	if err != nil {
		return models.ArchiveInfo{}, fmt.Errorf("не удалось прочитать архив: %v", err)
	}

	var files []models.FileInfo
	var totalSize float64
	for _, zipFile := range zipReader.File {
		size := float64(zipFile.UncompressedSize64)
		mimeType := utils.DetectMimeType(zipFile.Name)

		files = append(files, models.FileInfo{
			FilePath: zipFile.Name,
			Size:     size,
			MimeType: mimeType,
		})
		totalSize += size
	}

	return models.ArchiveInfo{
		Filename:    filename,
		ArchiveSize: float64(archiveSize),
		TotalSize:   totalSize,
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}

func CreateZipArchive(files []*multipart.FileHeader) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("ошибка открытия файла %s: %v", fileHeader.Filename, err)
		}
		defer file.Close()

		zipFileWriter, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			return nil, fmt.Errorf("ошибка создания архива для файла %s: %v", fileHeader.Filename, err)
		}

		if _, err = io.Copy(zipFileWriter, file); err != nil {
			return nil, fmt.Errorf("ошибка записи файла %s в архив: %v", fileHeader.Filename, err)
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, fmt.Errorf("ошибка закрытия архива: %v", err)
	}

	return buf.Bytes(), nil
}
