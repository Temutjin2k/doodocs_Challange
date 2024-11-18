package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/Temutjin2k/doodocs_Challange/internal/config"
	"github.com/Temutjin2k/doodocs_Challange/internal/service"
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
		SendError(w, "Unable to retrieve the file or the file not provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	mimeType := r.MultipartForm.File["file"][0].Header.Get("Content-Type")
	if !config.AvailiableMimeTypesToSendEmail[mimeType] {
		h.logger.Error("Invalid MimeType type", "MimeType", mimeType)
		SendError(w, fmt.Sprintf("Invalid MimeType: %v", mimeType), http.StatusBadRequest)
		return
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("Unable to read file", "Error", err)
		SendError(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Getting list of mails to send the file
	emails := r.MultipartForm.Value["emails"]

	// Validate email addresses
	h.logger.Info("Validating emails from request", "number", len(emails), "emails", emails)
	var emailList []string
	for _, email := range emails {
		if _, err := mail.ParseAddress(email); err != nil {
			h.logger.Error("Invalid email address", "email", email, "Error", err)
			SendError(w, fmt.Sprintf("Invalid email address: %s", email), http.StatusBadRequest)
			return
		}
		emailList = append(emailList, email)
	}
	h.logger.Info("Mails validated successfully", "number", len(emailList), "emails", emailList)

	err = h.mailService.SendFile(emailList, header.Filename, header.Header.Get("Content-Type"), fileData)
	if err != nil {
		h.logger.Error("Error sending emails", "Error", err)
		SendError(w, "Error sending emails", http.StatusInternalServerError)
		return
	}
}
