// Package service
package service

import (
	"github.com/n0mercy071/post-service-go/pkg/entity"
	"github.com/n0mercy071/post-service-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Post interface{}

type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
