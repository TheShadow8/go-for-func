package repository

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/TheShadow8/go-test-fiber/db"
	"gitlab.com/TheShadow8/go-test-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const FilesCollection = "files"

type FilesRepository interface {
	Save(files []*models.File) (*mongo.InsertManyResult, error)
	GetById(id string) (file *models.File, error error)
}

type filesRepository struct {
	c *mongo.Collection
}

func NewFilesRepository(conn db.Connection) FilesRepository {
	return &filesRepository{conn.DB().Collection(FilesCollection)}
}

func (r *filesRepository) Save(files []*models.File) (*mongo.InsertManyResult, error) {
	fmt.Println("files", files)
	r.c.InsertOne(context.TODO(), files[0])
	return r.c.InsertMany(context.TODO(), []interface{}{files})

}

func (r *filesRepository) GetById(id string) (file *models.File, erorr error) {
	fileID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	documentReturned := r.c.FindOne(context.TODO(), bson.M{"_id": fileID})

	fileDecode := models.File{}

	err = documentReturned.Decode(&fileDecode)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return &fileDecode, nil

}
