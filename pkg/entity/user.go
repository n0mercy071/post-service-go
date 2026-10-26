package entity

type User struct {
	ID           int    `json:"-" db:"id"`
	Username     string `json:"username" binding:"required" db:"username"`
	PasswordHash string `json:"password" binding:"required"`
}
