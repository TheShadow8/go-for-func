package services

import (
	"fmt"

	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileService interface {
	Save(files []*models.File) (*mongo.InsertManyResult, error)
	GetById(fileId string) (*models.File, error)
}

type fileServices struct {
	filesRepo repository.FilesRepository
}

func NewFileService(filesRepo repository.FilesRepository) FileService {
	return &fileServices{filesRepo}
}

func (s *fileServices) Save(files []*models.File) (*mongo.InsertManyResult, error) {

	insertResult, err := s.filesRepo.Save(files)

	fmt.Println("iR", insertResult)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (s *fileServices) GetById(fileId string) (*models.File, error) {
	return s.filesRepo.GetById(fileId)
}
