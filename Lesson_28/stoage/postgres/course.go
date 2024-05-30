

package postgres

import (
    "database/sql"
    "github.com/Go11Group/Dilfuza/Lesson_28/model"
)

type CourseRepo struct {
    DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
    return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {
    stmt, err := c.DB.Prepare("INSERT INTO courses (id, name) VALUES ($1, $2)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(course.ID, course.Name)
    if err != nil {
        return err
    }

    return nil
}

func (c *CourseRepo) GetByID(id string) (*model.Course, error) {
    var course model.Course
    err := c.DB.QueryRow("SELECT id, name FROM courses WHERE id = $1", id).Scan(&course.ID, &course.Name)
    if err != nil {
        return nil, err
    }
    return &course, nil
}

func (c *CourseRepo) Update(course *model.Course) error {
    stmt, err := c.DB.Prepare("UPDATE courses SET name = $1 WHERE id = $2")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(course.Name, course.ID)
    if err != nil {
        return err
    }

    return nil
}

func (c *CourseRepo) Delete(id string) error {
    stmt, err := c.DB.Prepare("DELETE FROM courses WHERE id = $1")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(id)
    if err != nil {
        return err
    }

    return nil
}