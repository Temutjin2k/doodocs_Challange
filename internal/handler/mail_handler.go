package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/mail"

	"github.com/Temutjin2k/doodocs_Challange/internal/service"
	"github.com/Temutjin2k/doodocs_Challange/utils"
)

type mailHandler struct {
	mailService service.MailServiceImpl
}

func NewMailHandler(mailService service.MailServiceImpl) *mailHandler {
	return &mailHandler{mailService: mailService}
}

func (h *mailHandler) SendMailHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !utils.IsValidMimeType(r.MultipartForm.File["file"][0]) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Getting list of mails to send the file
	emails := r.MultipartForm.Value

	// Validate email addresses
	emailList := []string{}

	for _, email := range emails["emails"] {
		if _, err := mail.ParseAddress(email); err != nil {
			http.Error(w, fmt.Sprintf("Invalid email address: %s", email), http.StatusBadRequest)
			return
		}
		emailList = append(emailList, email)
	}

	err = h.mailService.SendFile(emailList, header.Filename, header.Header.Get("Content-Type"), fileData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
