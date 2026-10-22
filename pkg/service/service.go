// Package service
package service

import "github.com/n0mercy071/post-service-go/pkg/repository"

type Authorization interface{}

type Post interface{}

type Service struct {
	Authorization
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
