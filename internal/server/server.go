package server

import (
	"net/http"
	"os"

	"github.com/Temutjin2k/doodocs_Challange/internal/handler"
	"github.com/Temutjin2k/doodocs_Challange/internal/logger"
	"github.com/Temutjin2k/doodocs_Challange/internal/middleware"
	"github.com/Temutjin2k/doodocs_Challange/internal/service"
)

func InitServer() http.Handler {
	// Logger
	logger := logger.InitLogger()

	// Services
	archServ := service.NewArchiveService()

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	mailServ, err := service.NewMailService(smtpHost, smtpPort, email, password)
	if err != nil {
		logger.Error("Mail Service creating", "Error", err)
		os.Exit(1)
	}

	// Handlers
	archHandler := handler.NewArchiveHandler(archServ)
	mailHandler := handler.NewMailHandler(mailServ)

	// Routes
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/archive/information", archHandler.ArchiveInformationHandler)
	mux.HandleFunc("POST /api/archive/files", archHandler.ArchiveFilesHandler)
	mux.HandleFunc("POST /api/mail/file", mailHandler.SendMailHandler)

	// Wrap the router with middlewares
	router := middleware.LoggingMiddleware(mux, logger)
	return router
}
