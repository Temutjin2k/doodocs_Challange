package server

import (
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/handler"
	"github.com/Temutjin2k/doodocs_Challange/internal/middleware"
	"github.com/Temutjin2k/doodocs_Challange/internal/service"
)

func InitServer() http.Handler {
	mux := http.NewServeMux()

	// Services
	archServ := service.NewArchiveService()
	mailServ := service.NewMailService()

	// Handlers
	archHandler := handler.NewArchiveHandler(archServ)
	mailHandler := handler.NewMailHandler(mailServ)

	// Routes
	mux.HandleFunc("/api/archive/information", archHandler.ArchiveInformationHandler)
	mux.HandleFunc("/api/archive/files", archHandler.ArchiveFilesHandler)
	mux.HandleFunc("/api/mail/file", mailHandler.SendMailHandler)

	// Wrap the router with middlewares
	router := middleware.LoggingMiddleware(mux)
	return router
}
