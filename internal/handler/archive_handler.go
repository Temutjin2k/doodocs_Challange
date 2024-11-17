package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/service"
)

type archiveHandler struct {
	archiveService service.ArchiveImpl
	logger         *slog.Logger
}

func NewArchiveHandler(archiveService service.ArchiveImpl, logger *slog.Logger) *archiveHandler {
	return &archiveHandler{
		archiveService: archiveService,
		logger:         logger,
	}
}

func (h *archiveHandler) ArchiveInformationHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		h.logger.Error("Unable to retrieve file from request", "Error", err)
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	h.logger.Info("Calling info method from archive service")
	archiveInfo, err := h.archiveService.Info(file, header)
	if err != nil {
		h.logger.Error("Info service error", "Error", err)
		http.Error(w, "Info service error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(archiveInfo)
}

func (h *archiveHandler) ArchiveFilesHandler(w http.ResponseWriter, r *http.Request) {
}
