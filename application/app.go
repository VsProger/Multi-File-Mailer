package application

import (
	"github.com/vsproger/Doodocs-days-2.0/config"
	"github.com/vsproger/Doodocs-days-2.0/logger"
	"net/http"
)

type Application struct {
	Config *config.Config
	Logger *logger.Logger
}

func NewApplication(cfg *config.Config, log *logger.Logger) *Application {
	return &Application{
		Config: cfg,
		Logger: log,
	}
}

func (app *Application) SetupRoutes() {
	http.HandleFunc("/api/archive/information", app.ArchiveInfoHandler)
	http.HandleFunc("/api/archive/files", app.CreateArchiveHandler)
	http.HandleFunc("/api/mail/file", app.SendMailHandler)
}

func (app *Application) Start() {
	app.Logger.Info("Запуск сервера на порту :8080")
	http.ListenAndServe(":8080", nil)
}
