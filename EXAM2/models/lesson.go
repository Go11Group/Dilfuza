package models

import "time"

type Lesson struct {
	LessonID   string    `json:"lesson_id" db:"lesson_id"`
	CourseID   string    `json:"course_id" db:"course_id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt  int64     `json:"deleted_at" db:"deleted_at"`
}