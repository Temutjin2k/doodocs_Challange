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
		SendError(w, "Unable to retrieve file or file not provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	h.logger.Info("Calling info method from archive service")
	archiveInfo, err := h.archiveService.Info(file, header)
	if err != nil {
		h.logger.Error("Info service error", "Error", err)
		SendError(w, "Info service error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(archiveInfo)
	if err != nil {
		SendError(w, "Could not encode the reponse", http.StatusInternalServerError)
	}
}

func (h *archiveHandler) ArchiveFilesHandler(w http.ResponseWriter, r *http.Request) {
}
