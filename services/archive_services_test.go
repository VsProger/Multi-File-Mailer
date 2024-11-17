package services

import (
	"archive/zip"
	"bytes"
	"github.com/vsproger/Doodocs-days-2.0/models"
	"github.com/vsproger/Doodocs-days-2.0/utils"
	"io"
	"testing"
)

type testMultipartFile struct {
	*io.SectionReader
}

func (f *testMultipartFile) Close() error {
	return nil
}

func TestProcessArchive(t *testing.T) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	files := []struct {
		Name    string
		Content string
	}{
		{"file1.txt", "Содержимое файла 1"},
		{"file2.txt", "Содержимое файла 2"},
	}

	var expectedFiles []models.FileInfo
	var expectedTotalSize float64

	for _, file := range files {
		f, err := zipWriter.Create(file.Name)
		if err != nil {
			t.Fatalf("Ошибка при добавлении файла в архив: %v", err)
		}
		_, err = f.Write([]byte(file.Content))
		if err != nil {
			t.Fatalf("Ошибка при записи содержимого файла: %v", err)
		}

		size := float64(len(file.Content))
		mimeType := utils.DetectMimeType(file.Name)
		expectedFiles = append(expectedFiles, models.FileInfo{
			FilePath: file.Name,
			Size:     size,
			MimeType: mimeType,
		})
		expectedTotalSize += size
	}

	if err := zipWriter.Close(); err != nil {
		t.Fatalf("Ошибка при закрытии архива: %v", err)
	}

	fileReader := bytes.NewReader(buf.Bytes())
	archiveSize := int64(fileReader.Len())
	file := &testMultipartFile{io.NewSectionReader(fileReader, 0, archiveSize)}

	archiveInfo, err := ProcessArchive(file, "test.zip", archiveSize)
	if err != nil {
		t.Fatalf("ProcessArchive вернула ошибку: %v", err)
	}

	if archiveInfo.Filename != "test.zip" {
		t.Errorf("Ожидаемое имя файла 'test.zip', получили '%s'", archiveInfo.Filename)
	}

	if archiveInfo.ArchiveSize != float64(archiveSize) {
		t.Errorf("Ожидаемый размер архива %f, получили %f", float64(archiveSize), archiveInfo.ArchiveSize)
	}

	if archiveInfo.TotalSize != expectedTotalSize {
		t.Errorf("Ожидаемый общий размер файлов %f, получили %f", expectedTotalSize, archiveInfo.TotalSize)
	}

	if archiveInfo.TotalFiles != float64(len(expectedFiles)) {
		t.Errorf("Ожидаемое количество файлов %d, получили %f", len(expectedFiles), archiveInfo.TotalFiles)
	}

	for i, fileInfo := range archiveInfo.Files {
		expectedFile := expectedFiles[i]
		if fileInfo.FilePath != expectedFile.FilePath {
			t.Errorf("Ожидаемый путь '%s', получили '%s'", expectedFile.FilePath, fileInfo.FilePath)
		}
		if fileInfo.Size != expectedFile.Size {
			t.Errorf("Ожидаемый размер файла %f, получили %f", expectedFile.Size, fileInfo.Size)
		}
		if fileInfo.MimeType != expectedFile.MimeType {
			t.Errorf("Ожидаемый MIME-тип '%s', получили '%s'", expectedFile.MimeType, fileInfo.MimeType)
		}
	}
}
