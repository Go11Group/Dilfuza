package models


type Enrollments struct {
    EnrollmentId     string `json:"enrollment_id"db:"enrollment_id"`
    UserId           string `json:"user_id" db:"user_id"`
    CourseId         string `json:"course_id" db:"course_id"`
    EnrollmentDate  string `json:"enrollment_date" db:"enrollment_date"`
    CreatedAt    	string `json:"created_at" db:"created_at"`
    UpdatedAt       string `json:"updated_at" db:"updated_at"`
    DeletedAt      string `json:"deleted_at" db:"deleted_at"`
}

