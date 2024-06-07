package model

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	Id              uuid.UUID `json:"id"`
	CourseName  string    `json:"course_name"`
	Rector          string    `json:"rector"`
	CourseEmail string    `json:"course_email"`
	CoursePhone string    `json:"course_phone"`
	CreatedAt       time.Time `json:"created_at"`
	DeletedAt       time.Time `json:"deleted_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
