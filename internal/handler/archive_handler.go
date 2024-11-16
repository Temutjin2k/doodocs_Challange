package handler

import (
	"encoding/json"
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
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	archiveInfo, err := h.archiveService.Info(file, header)
	if err != nil {
		http.Error(w, "Info service error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(archiveInfo)
}

func (h *archiveHandler) ArchiveFilesHandler(w http.ResponseWriter, r *http.Request) {
}
