package entity

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"titie"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PostTag struct {
	ID     int `json:"id"`
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}

type PostAttacment struct {
	ID       int    `json:"id"`
	PostID   int    `json:"post_id"`
	FilePath string `json:"file_path"`
}
