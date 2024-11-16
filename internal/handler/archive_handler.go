package handler

import (
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/service"
)

type archiveHandler struct {
	archiveService service.ArchiveImpl
}

func NewArchiveHandler(archiveService service.ArchiveImpl) *archiveHandler {
	return &archiveHandler{archiveService: archiveService}
}

func (h *archiveHandler) ArchiveInformationHandler(w http.ResponseWriter, r *http.Request) {
}

func (h *archiveHandler) ArchiveFilesHandler(w http.ResponseWriter, r *http.Request) {
}
