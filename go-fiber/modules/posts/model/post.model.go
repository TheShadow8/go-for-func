package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt *time.Time         `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
}
