package entity

type User struct {
	ID           int    `json:"-"`
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
}
