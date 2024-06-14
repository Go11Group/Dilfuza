package models

import "time"

type Enrollment struct {
	EnrollmentID   string	 `json:"enrollment_id" db:"enrollment_id"`
	UserID         string    `json:"user_id" db:"user_id"`
	CourseID       string    `json:"course_id" db:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date" db:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      int64     `json:"deleted_at" db:"deleted_at"`
}