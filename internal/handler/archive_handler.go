package handler

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/Temutjin2k/doodocs_Challange/config"
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
	err := r.ParseMultipartForm(1 << 30) // 1 GB
	if err != nil {
		h.logger.Error("Could not parse MultipartForm", "Error", err)
		SendError(w, "Could not parse form data", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files[]"]

	if len(files) == 0 {
		h.logger.Error("No files send to archive")
		SendError(w, "No files in request", http.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileHeader := range files {
		// Validate for MimeType
		if config.AvailibleMimeTypesToArvhive[fileHeader.Header.Get("Content-Type")] {
		}

		file, err := fileHeader.Open()
		if err != nil {
			h.logger.Error("Unable to open file", "Error", err)
			SendError(w, "Unable to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Сreating file in archvie
		zipFile, err := zipWriter.Create(filepath.Base(fileHeader.Filename))
		if err != nil {
			h.logger.Error("Unable to create zip entry", "Error", err)
			SendError(w, "Unable to create zip entry", http.StatusInternalServerError)
			return
		}

		// Copying file data to archive
		if _, err := io.Copy(zipFile, file); err != nil {
			h.logger.Error("Unable to write file to zip", "Error", err)
			SendError(w, "Unable to write file to zip", http.StatusInternalServerError)
			return
		}
	}

	if err := zipWriter.Close(); err != nil {
		h.logger.Error("Unable to close zip file", "Error", err)
		SendError(w, "Unable to close zip file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", `attachment; filename="archive.zip"`)

	if _, err := w.Write(buf.Bytes()); err != nil {
		h.logger.Error("Zip file was created, but couldn't send it to client", "Error", err)
		SendError(w, "Unable to send zip file", http.StatusInternalServerError)
	}
}
