package models

import "time"

type Course struct {
	CourseID    string    `json:"course_id" db:"course_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   int64     `json:"deleted_at" db:"deleted_at"`
}