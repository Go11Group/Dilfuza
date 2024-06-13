package postgres

import (
	"database/sql"
	"github.com/Go11Group/Dilfuza/EXAM2/models"
	"time"
)

type EnrollmentRepo struct {
	db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}


func (e *EnrollmentRepo) EnrollmentGet(id string) (models.Enrollments, error) {
	var enrollment models.Enrollments
	err := e.db.QueryRow("SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments WHERE enrollment_id = $1", id).Scan(
		&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	return enrollment, err
}


func (e *EnrollmentRepo) EnrollmentCreate(enrollment models.Enrollments) error {
	query := `INSERT INTO enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := e.db.Exec(query, enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId, time.Now(), time.Now(), enrollment.DeletedAt)
	return err
}
func (e *EnrollmentRepo) EnrollmentUpdate(enrollment models.Enrollments) error {
	query := `UPDATE enrollments SET user_id = $1, course_id = $2, enrollment_date = $3, updated_at = $4, deleted_at = $5 WHERE enrollment_id = $6`
	_, err := e.db.Exec(query, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, time.Now(), enrollment.DeletedAt, enrollment.EnrollmentId)
	return err
}


func (e *EnrollmentRepo) EnrollmentDelete(id string) error {
	query := `UPDATE enrollments SET deleted_at = $1 WHERE enrollment_id = $2`
	_, err := e.db.Exec(query, time.Now(), id)
	return err
}

