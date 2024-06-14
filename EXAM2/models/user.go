package models

import (
	"time"
)
type User struct {
	UserID    string    `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Birthday  time.Time `json:"birthday" db:"birthday"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt int64     `json:"deleted_at" db:"deleted_at"`
}