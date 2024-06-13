package postgres

import (
	"database/sql"
	"time"

	"github.com/Go11Group/Dilfuza/EXAM2/models"
	_ "github.com/lib/pq"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (c *CourseRepo) CourseGet(id string) (models.Courses, error) {
	var course models.Courses
	err := c.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1", id).Scan(
		&course.CourseId, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
	return course, err
}

func (c *CourseRepo) CourseCreate(course models.Courses) error {
	query := `INSERT INTO courses (course_id, title, description, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := c.db.Exec(query, course.CourseId, course.Title, course.Description, time.Now(), time.Now(), nil)
	return err
}

func (c *CourseRepo) CourseUpdate(course models.Courses) error {
	query := `UPDATE courses SET title = $1, description = $2, updated_at = $3, deleted_at = $4 WHERE course_id = $5`
	_, err := c.db.Exec(query, course.Title, course.Description, time.Now(), course.DeletedAt, course.CourseId)
	return err
}

func (c *CourseRepo) CourseDelete(id string) error {
	query := `UPDATE courses SET deleted_at = $1 WHERE course_id = $2`
	_, err := c.db.Exec(query, time.Now(), id)
	return err
}
