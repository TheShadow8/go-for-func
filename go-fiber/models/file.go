package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Project   string             `json:"project" bson:"project"`
	FileName  string             `json:"file_name" bson:"file_name"`
	FilePath  string             `json:"file_path" bson:"file_path"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
