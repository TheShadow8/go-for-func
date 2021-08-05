package posts

import (
	"context"

	"gitlab.com/TheShadow8/go-test-fiber/db"
	posts "gitlab.com/TheShadow8/go-test-fiber/modules/posts/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const PostCollection = "posts"

type PostRepository interface {
	Save(post *posts.Post) (*mongo.InsertOneResult, error)
	GetAll() ([]*posts.Post, error)
}

type postRepository struct {
	c *mongo.Collection
}

func NewPostRepository(conn db.Connection) PostRepository {
	return &postRepository{conn.DB().Collection(PostCollection)}
}

func (r *postRepository) Save(post *posts.Post) (*mongo.InsertOneResult, error) {
	return r.c.InsertOne(context.TODO(), post)
}

func (r *postRepository) GetAll() (posts []*posts.Post, err error) {
	cursor, err := r.c.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if cursor.All(context.TODO(), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
