package service

import "github.com/Temutjin2k/doodocs_Challange/models"

type ArchiveImpl interface {
	Info() (models.ArchiveFile, error)
	ArchiveFiles() error
}

type archiveService struct{}

func NewArchiveService() *archiveService {
	return &archiveService{}
}

func (s *archiveService) Info() (models.ArchiveFile, error) {
	return models.ArchiveFile{}, nil
}

func (s *archiveService) ArchiveFiles() error {
	return nil
}
