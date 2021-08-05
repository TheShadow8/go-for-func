package posts

import (
	posts "gitlab.com/TheShadow8/go-test-fiber/modules/posts/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService interface {
	Save(post *posts.Post) (*mongo.InsertOneResult, error)
	GetAll() ([]*posts.Post, error)
}

type postService struct {
	postsRepo PostRepository
}

func NewPostSevice(postRepo PostRepository) PostService {
	return &postService{postRepo}
}

func (s *postService) Save(post *posts.Post) (*mongo.InsertOneResult, error) {
	return s.postsRepo.Save(post)
}

func (s *postService) GetAll() ([]*posts.Post, error) {
	return s.postsRepo.GetAll()
}
