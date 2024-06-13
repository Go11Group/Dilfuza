package postgres

import (
	"database/sql"
	"github.com/Go11Group/Dilfuza/EXAM2/models"
	"time"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}


func (l *LessonRepo) Lesson_Get(id string) (models.Lessons, error) {
	var lesson models.Lessons
	err := l.db.QueryRow("SELECT lesson_id, course_id, title, content, created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1", id).Scan(
		&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	return lesson, err
}


func (l *LessonRepo) Lesson_Create(lesson models.Lessons) error {
	query := `INSERT INTO lessons (lesson_id, course_id, title, content, created_at, updated_at, deleted_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := l.db.Exec(query, lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), time.Now(), nil)
	return err
}


func (l *LessonRepo) Lesson_Update(lesson models.Lessons) error {
	query := `UPDATE lessons SET course_id = $1, title = $2, content = $3, updated_at = $4, deleted_at = $5 WHERE lesson_id = $6`
	_, err := l.db.Exec(query, lesson.CourseId, lesson.Title, lesson.Content, time.Now(), lesson.DeletedAt, lesson.LessonId)
	return err
}


func (l *LessonRepo) Lesson_Delete(id string) error {
	query := `UPDATE lessons SET deleted_at = $1 WHERE lesson_id = $2`
	_, err := l.db.Exec(query, time.Now(), id)
	return err
}

