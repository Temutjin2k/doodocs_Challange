package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Temutjin2k/doodocs_Challange/internal/config"
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
		h.logger.Error("Could not retrieve file from request", "Error", err)
		SendError(w, "Could not retrieve file or file not provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if header.Size > config.MaxInfoArchiveSize {
		h.logger.Error("Too large file", "Filename", header.Filename, "Size", header.Size)
		SendError(w, "Could not get information from archive. File size is too large", http.StatusBadRequest)
		return
	}

	h.logger.Info("Calling info method from archive service")
	archiveInfo, err := h.archiveService.Info(file, header)
	if err != nil {
		h.logger.Error("Info service error", "Error", err)
		SendError(w, "Could not get information from archive", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(archiveInfo)
	if err != nil {
		h.logger.Error("Could not encode the reponse", "Error", err)
		SendError(w, "Could not encode the reponse", http.StatusInternalServerError)
	}
}

func (h *archiveHandler) ArchiveFilesHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(config.MaxArchivingSize) // 10 GB
	if err != nil {
		h.logger.Error("Could not parse MultipartForm", "Error", err)
		SendError(w, "Could not parse form data", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files[]"]

	archiveData, err := h.archiveService.ArchiveFiles(files)
	if err != nil {
		h.logger.Error("Could not create archive", "Error", err)
		SendError(w, "Could not create archive", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="archive.zip"`)

	if _, err := w.Write(archiveData); err != nil {
		h.logger.Error("Zip file was created, but couldn't send it to client", "Error", err)
		SendError(w, "Could not send zip file", http.StatusInternalServerError)
	}
}
