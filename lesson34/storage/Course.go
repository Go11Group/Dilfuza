package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Dilfuza/lesson34/model"
)

type CourseRepository struct {
	Db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		Db: db,
	}
}

func (un *CourseRepository) CreateCourse(course model.Course) error {
	_, err := un.Db.Exec("insert into  courses(coursename,rector,courseemail) values ($1,$2,$3)", course.CourseName, course.Rector, course.CourseEmail)
	if err != nil {
		return err
	}
	return nil
}

func (un *CourseRepository) UpdateCourse(id string, course model.Course) error {
	_, err := un.Db.Exec("update courses's set coursename=$1,rector=$2,courseemail=$3 where id=$4", course.CourseName, course.Rector, course.CourseEmail, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepository) DeleteCourse(id string) error {
	_, err := c.Db.Exec("delete  from  courses where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (un *CourseRepository) ReadAllCourse() ([]model.Course, error) {
	rows, err := un.Db.Query("select id , coursename,rector, courseemail ,created_at,updated_at,deleted_at  from courses")
	courses := []model.Course{}
	if err != nil {
		fmt.Println("golang", err)
		return nil, err
	}
	course := model.Course{}
	for rows.Next() {
		err := rows.Scan(&course.Id, &course.CourseName, &course.Rector, &course.CourseEmail, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)

		if err != nil {
			fmt.Println("go11", err)
			return nil, err
		}

		courses = append(courses, course)
	}
	fmt.Println("golang", courses)

	return courses, err
}
