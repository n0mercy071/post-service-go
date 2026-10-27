package service

import (
	"github.com/n0mercy071/post-service-go/pkg/entity"
	"github.com/n0mercy071/post-service-go/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(userID int, post entity.Post) (entity.Post, error) {
	post.UserID = userID

	return s.repo.CreatePost(post)
}
