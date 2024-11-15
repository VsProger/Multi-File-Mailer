package application

import (
	"encoding/json"
	"fmt"
	"github.com/vsproger/Doodocs-days-2.0/services"
	"github.com/vsproger/Doodocs-days-2.0/utils"
	"net/http"
)

// ArchiveInfoHandler обрабатывает запрос информации об архиве
func (app *Application) ArchiveInfoHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		app.Logger.Error("Не удалось получить файл из запроса")
		http.Error(w, "Не удалось получить файл из запроса", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//if !utils.IsValidArchiveMimeType(fileHeader) {
	//	app.Logger.Error("Недопустимый тип файла для архива")
	//	http.Error(w, "Недопустимый тип файла", http.StatusUnsupportedMediaType)
	//	return
	//}

	info, err := services.ProcessArchive(file, fileHeader.Filename, fileHeader.Size)
	if err != nil {
		app.Logger.Error(fmt.Sprintf("Ошибка обработки архива: %v", err))
		http.Error(w, fmt.Sprintf("Ошибка обработки архива: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(info)
	if err != nil {
		app.Logger.Error("Не удалось закодировать информацию об архиве в JSON")
		http.Error(w, "Ошибка сервера: не удалось закодировать информацию об архиве", http.StatusInternalServerError)
		return
	}

	app.Logger.Info("Архив обработан успешно")
}

// CreateArchiveHandler обрабатывает создание архива
func (app *Application) CreateArchiveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		app.Logger.Error("Не удалось разобрать multipart form: " + err.Error())
		http.Error(w, "Ошибка сервера: не удалось обработать форму", http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["files[]"]
	if len(files) == 0 {
		app.Logger.Error("Не передано ни одного файла для создания архива")
		http.Error(w, "Не передано ни одного файла", http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {
		if !utils.IsValidArchiveMimeType(fileHeader) {
			app.Logger.Error(fmt.Sprintf("Недопустимый MIME-тип для файла %s", fileHeader.Filename))
			http.Error(w, fmt.Sprintf("Файл %s имеет недопустимый MIME-тип", fileHeader.Filename), http.StatusUnsupportedMediaType)
			return
		}
	}

	zipData, err := services.CreateZipArchive(files)
	if err != nil {
		app.Logger.Error(fmt.Sprintf("Ошибка создания архива: %v", err))
		http.Error(w, fmt.Sprintf("Ошибка создания архива: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\"archive.zip\"")
	_, err = w.Write(zipData)
	if err != nil {
		app.Logger.Error("Не удалось отправить ZIP-файл клиенту: " + err.Error())
		http.Error(w, "Ошибка сервера: не удалось отправить файл", http.StatusInternalServerError)
		return
	}
	app.Logger.Info("Архив успешно создан и отправлен")
}
