package models

type Courses struct{ 
    CourseId        string `json:"course_id" db:"course_id"`
    Title           string `json:"title" db:"title"` 
    Description     string `json:"description" db:"description"`
	CreatedAt    	string `json:"created_at" db:"created_at"`
    UpdatedAt       string `json:"updated_at" db:"updated_at"`
    DeletedAt      string `json:"deleted_at" db:"deleted_at"`

}


