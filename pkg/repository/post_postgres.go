package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/n0mercy071/post-service-go/pkg/entity"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{
		db: db,
	}
}

func (r *PostPostgres) CreatePost(post entity.Post) (entity.Post, error) {
	fail := func(err error) (entity.Post, error) {
		return entity.Post{}, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fail(err)
	}

	defer tx.Rollback()

	if err := r.insertPostAndUpdateStruct(tx, &post); err != nil {
		fail(err)
	}

	if err := tx.Commit(); err != nil {
		fail(err)
	}

	return post, nil
}

func (r *PostPostgres) insertPostAndUpdateStruct(tx *sql.Tx, post *entity.Post) error {
	var postID int
	query := fmt.Sprintf("INSERT INTO %s (title, user_id) values ($1, $2) RETURNING id", postsTable)
	row := tx.QueryRow(query, post.Title, post.UserID)
	if err := row.Scan(&postID); err != nil {
		return err
	}

	post.ID = postID

	return nil
}
