package repository

import (
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

func (r *PostPostgres) CreatePost(post entity.Post) (int, error) {
	fail := func(err error) (int, error) {
		return 0, err
	}

	tx := r.db.MustBegin()

	defer tx.Rollback()

	postID, err := r.insertPost(tx, post)
	if err != nil {
		return fail(err)
	}
	err = r.insertTags(tx, postID, post.Tags)
	if err != nil {
		return fail(err)
	}

	if err := tx.Commit(); err != nil {
		return fail(err)
	}

	return postID, nil
}

func (r *PostPostgres) insertPost(tx *sqlx.Tx, post entity.Post) (int, error) {
	var postID int
	query := fmt.Sprintf(
		"INSERT INTO %s (title, user_id) values ($1, $2) RETURNING id",
		postsTable,
	)
	row := tx.QueryRow(query, post.Title, post.UserID)
	if err := row.Scan(&postID); err != nil {
		return 0, err
	}

	return postID, nil
}

func (r *PostPostgres) insertTags(tx *sqlx.Tx, postID int, tags []entity.Tag) error {
	for _, tag := range tags {
		_, err := tx.Exec(
			`INSERT INTO tags (name, user_id) values ($1, $2)
							ON CONFLICT (name) DO NOTHING`,
			tag.Name,
			tag.UserID,
		)
		if err != nil {
			return err
		}

		var tagID int
		err = tx.Get(
			&tagID,
			"SELECT id FROM tags WHERE name = $1",
			tag.Name,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			`INSERT INTO posts_tags (post_id, tag_id) values ($1, $2)
							ON CONFLICT DO NOTHING`,
			postID,
			tagID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PostPostgres) GetByID(postID int) (entity.Post, error) {
	var post entity.Post

	err := r.db.Get(
		&post,
		"SELECT id, title, user_id FROM posts WHERE id = $1",
		postID,
	)
	if err != nil {
		return post, err
	}

	err = r.db.Select(
		&post.Tags,
		`SELECT t.id, t.name, t.user_id
						FROM tags t
							JOIN posts_tags pt ON pt.tag_id = t.id
						WHERE pt.post_id = $1`,
		post.ID,
	)
	if err != nil {
		return post, err
	}

	err = r.db.Select(
		&post.Attachments,
		`SELECT id, post_id, user_id, file_path
						FROM post_attachments
						WHERE post_id = $1`,
		post.ID,
	)

	return post, err
}
