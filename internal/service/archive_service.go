package service

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/Temutjin2k/doodocs_Challange/internal/config"
	"github.com/Temutjin2k/doodocs_Challange/models"
)

type ArchiveImpl interface {
	Info(multipart.File, *multipart.FileHeader) (models.ArchiveFile, error)
	ArchiveFiles([]*multipart.FileHeader) ([]byte, error)
}

type archiveService struct{}

func NewArchiveService() *archiveService {
	return &archiveService{}
}

func (s *archiveService) Info(file multipart.File, header *multipart.FileHeader) (models.ArchiveFile, error) {
	// Temp file to use zip package
	tempFile, err := os.CreateTemp("", "uploaded-*.zip")
	if err != nil {
		return models.ArchiveFile{}, err
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.ReadFrom(file)
	if err != nil {
		return models.ArchiveFile{}, err
	}
	tempFile.Close()

	// Open the ZIP archive
	archive, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return models.ArchiveFile{}, err
	}
	defer archive.Close()

	var files []models.File
	var totalSize float64
	for _, f := range archive.File {
		if !f.FileInfo().IsDir() {
			mimeType := mime.TypeByExtension(filepath.Ext(f.Name))
			files = append(files, models.File{
				File_path: f.Name,
				Size:      float64(f.FileInfo().Size()),
				Mimetype:  mimeType,
			})
			totalSize += float64(f.FileInfo().Size())
		}
	}

	// Prepare the response
	response := models.ArchiveFile{
		Filename:     header.Filename,
		Archive_size: float64(header.Size),
		Totalsize:    float64(totalSize),
		Total_files:  len(files),
		Files:        files,
	}

	return response, nil
}

func (s *archiveService) ArchiveFiles(files []*multipart.FileHeader) ([]byte, error) {
	if len(files) == 0 {
		return []byte{}, errors.New("no file to archvie")
	}

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileHeader := range files {
		// Validate for MimeType
		mimeType := fileHeader.Header.Get("Content-Type")
		if !config.AvailiableMimeTypesToArvhive[mimeType] {
			return []byte{}, fmt.Errorf("MimeType not allowed. Filename: %v, MimeType: %v", fileHeader.Filename, mimeType)
		}

		file, err := fileHeader.Open()
		if err != nil {
			return []byte{}, errors.New("unable to open " + fileHeader.Filename)
		}
		defer file.Close()

		// Сreating file in archvie
		zipFile, err := zipWriter.Create(filepath.Base(fileHeader.Filename))
		if err != nil {
			return []byte{}, errors.New("unable to create zip entry. Filename: " + fileHeader.Filename)
		}

		// Copying file data to archive
		if _, err := io.Copy(zipFile, file); err != nil {
			return []byte{}, errors.New("unable to write file to zip. Filename: " + fileHeader.Filename)
		}
	}

	if err := zipWriter.Close(); err != nil {
		return []byte{}, errors.New("unable to close zip file")
	}

	return buf.Bytes(), nil
}
