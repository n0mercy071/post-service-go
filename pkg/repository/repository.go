package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, passwordHash string) (entity.User, error)
}

type Post interface {
	CreatePost(post entity.Post) (entity.Post, error)
}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
