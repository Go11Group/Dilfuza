package models

import "time"

type  Lessons struct {
    LessonId string `json:"lesson_id" db:"lesson_id"`
    CourseId string  `json:"course_id" db:"course_id"`
    Title     string  `json:"title" db:"title"`  
    Content   string  `json:"content" db:"content"`
    CreatedAt string  `json:"created_at" db:"created_at"`
    UpdatedAt string  `json:"updated_at" db:"updated_at"`
    DeletedAt  *time.Time  `json:"deleted_at" db:"deleted_at"`
}


