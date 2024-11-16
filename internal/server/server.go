package server

import (
	"net/http"

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
	mailServ := service.NewMailService()

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
