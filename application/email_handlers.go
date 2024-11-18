package application

import (
	"github.com/vsproger/Doodocs-days-2.0/services"
	"github.com/vsproger/Doodocs-days-2.0/utils"
	"net/http"
)

func (app *Application) SendMailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		app.Logger.Error("Не удалось получить файл из запроса")
		http.Error(w, "Не удалось получить файл из запроса", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !utils.IsValidEmailMimeType(fileHeader) {
		app.Logger.Error("Недопустимый тип файла для отправки по почте")
		http.Error(w, "Недопустимый тип файла для отправки", http.StatusUnsupportedMediaType)
		return
	}

	emails := r.FormValue("emails")
	if emails == "" {
		app.Logger.Error("Не указан список почт")
		http.Error(w, "Не указан список почт", http.StatusBadRequest)
		return
	}

	emailList := utils.ParseEmails(emails)
	err = services.SendEmailWithAttachment(emailList, file, fileHeader.Filename, fileHeader.Header.Get("Content-Type"), app.Config)
	if err != nil {
		app.Logger.Error("Не удалось отправить файл по электронной почте: " + err.Error())
		http.Error(w, "Ошибка сервера: не удалось отправить файл по электронной почте", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Файл успешно отправлен на указанные почты"))
	if err != nil {
		app.Logger.Error("Не удалось отправить ответ клиенту: " + err.Error())
		http.Error(w, "Ошибка сервера: не удалось отправить подтверждение", http.StatusInternalServerError)
		return
	}
	app.Logger.Info("Файл успешно отправлен по почте")
}
