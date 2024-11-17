package application

import (
	"github.com/vsproger/Doodocs-days-2.0/config"
	"github.com/vsproger/Doodocs-days-2.0/logger"
	"net/http"
	"os"
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

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/archive/information", app.ArchiveInfoHandler)
	mux.HandleFunc("/api/archive/files", app.CreateArchiveHandler)
	mux.HandleFunc("/api/mail/file", app.SendMailHandler)
	return mux
}

func (app *Application) Start() {
	app.Logger.Info("Запуск сервера на порту " + app.Config.Port)
	err := http.ListenAndServe(app.Config.Port, app.Routes())
	app.Logger.Error(err.Error())
	os.Exit(1)
}
