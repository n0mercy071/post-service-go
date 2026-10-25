// Package repository
package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type Post interface{}

type Repository struct {
	Authorization
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
