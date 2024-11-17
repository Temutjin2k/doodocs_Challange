package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/Temutjin2k/doodocs_Challange/internal/service"
	"github.com/Temutjin2k/doodocs_Challange/utils"
)

type mailHandler struct {
	mailService service.MailServiceImpl
	logger      *slog.Logger
}

func NewMailHandler(mailService service.MailServiceImpl, logger *slog.Logger) *mailHandler {
	return &mailHandler{mailService: mailService, logger: logger}
}

func (h *mailHandler) SendMailHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		h.logger.Error("Unable to get file from form", "Error", err)
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !utils.IsValidMimeType(r.MultipartForm.File["file"][0]) {
		h.logger.Error("Invalid MimeType type", "Error", err)
		http.Error(w, "Invalid MimeType type", http.StatusBadRequest)
		return
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("Unable to read file", "Error", err)
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Getting list of mails to send the file
	emails := r.MultipartForm.Value

	// Validate email addresses
	h.logger.Info("Validating emails from request")

	emailList := []string{}
	for _, email := range emails["emails"] {
		if _, err := mail.ParseAddress(email); err != nil {
			h.logger.Error("Invalid email address", "email", email, "Error", err)
			http.Error(w, fmt.Sprintf("Invalid email address: %s", email), http.StatusBadRequest)
			return
		}
		emailList = append(emailList, email)
	}
	h.logger.Info("Mails validated successful", "number", len(emailList), "emails", emailList)

	err = h.mailService.SendFile(emailList, header.Filename, header.Header.Get("Content-Type"), fileData)
	if err != nil {
		h.logger.Error("Error sending emails", "Error", err)
		http.Error(w, "Error sending emails", http.StatusInternalServerError)
		return
	}
}
