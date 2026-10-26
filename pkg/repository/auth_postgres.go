package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, passwordHash string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
