package services

import (
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/repository"
)

type FileServices interface {
	Save(files []*models.File) (*mongo.InsertManyResult, error)
	GetById(fileId string) (*models.File, error)
}

type fileServices struct {
	filesRepo repository.FilesRepository
}

func NewFileService(filesRepo repository.FilesRepository) FileServices {
	return &fileServices{filesRepo}
}

func (s *fileServices) Save(files []*models.File) (*mongo.InsertManyResult, error) {

	insertedResults, err := s.filesRepo.Save(files)

	if err != nil {
		return nil, err
	}

	return insertedResults, nil
}

func (s *fileServices) GetById(fileId string) (*models.File, error) {
	return s.filesRepo.GetById(fileId)
}
