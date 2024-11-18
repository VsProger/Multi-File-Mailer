package utils

import (
	"fmt"
	"github.com/h2non/filetype"
	"io"
	"mime/multipart"
)

var supportedArchives = map[string]bool{
	"zip": true,
	"tar": true,
	"gz":  true,
	"bz2": true,
	"7z":  true,
	"rar": true,
	"ar":  true,
}

func IsArchive(file multipart.File) (bool, error) {
	buffer := make([]byte, 261)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return false, fmt.Errorf("невозможно прочитать файл: %v", err)
	}
	if n < 4 {
		return false, fmt.Errorf("файл слишком мал для проверки сигнатуры")
	}

	kind, unknown := filetype.Match(buffer)
	if unknown != nil {
		return false, fmt.Errorf("не удалось определить тип файла: %v", unknown)
	}

	fmt.Printf("Определенный тип файла: %s\n", kind.MIME.Value)

	isSupported := supportedArchives[kind.Extension]
	return isSupported, nil
}
