package handler

import (
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/service"
)

type mailHandler struct {
	mailService service.MailServiceImpl
}

func NewMailHandler(mailService service.MailServiceImpl) *mailHandler {
	return &mailHandler{mailService: mailService}
}

func (h *mailHandler) SendMailHandler(w http.ResponseWriter, r *http.Request) {
}
