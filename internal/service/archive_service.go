package service

import (
	"archive/zip"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/Temutjin2k/doodocs_Challange/models"
)

type ArchiveImpl interface {
	Info(multipart.File, *multipart.FileHeader) (models.ArchiveFile, error)
	ArchiveFiles() error
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
		if f.FileInfo().IsDir() {
			files = append(files, models.File{
				File_path: f.Name,
				Size:      0,
				Mimetype:  "directory",
			})
		} else {
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

func (s *archiveService) ArchiveFiles() error {
	return nil
}
