package entity

type Post struct {
	ID         int              `json:"id" db:"id"`
	Title      string           `json:"title" db:"title" binding:"required"`
	UserID     int              `json:"user_id" db:"user_id"`
	Tags       []Tag            `json:"tags"`
	Attachment []PostAttachment `json:"attachments"`
}

type Tag struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name" binding:"required"`
	UserID int    `json:"user_id" db:"user_id"`
}

type PostTag struct {
	PostID int `json:"post_id" db:"post_id"`
	TagID  int `json:"tag_id" db:"tag_id"`
}

type PostAttachment struct {
	ID       int    `json:"id" db:"id"`
	PostID   int    `json:"post_id" db:"post_id"`
	FilePath string `json:"file_path" db:"file_path"`
	UserID   int    `json:"user_id" db:"user_id"`
}
