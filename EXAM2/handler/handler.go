package handler

import "github.com/Go11Group/Dilfuza/EXAM2/storage/postgres"

type Handler struct {
    Users       *postgres.UserRepo
    Courses     *postgres.CourseRepo
    Lessons     *postgres.LessonRepo
    Enrollments *postgres.EnrollmentRepo
}
