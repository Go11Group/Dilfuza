package main

import (
    "log"

    "github.com/Go11Group/Dilfuza/EXAM2/storage/postgres"
    "github.com/Go11Group/Dilfuza/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    db, err := postgres.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    userRepo := postgres.NewUserRepo(db)
    courseRepo := postgres.NewCourseRepo(db)
    lessonRepo := postgres.NewLessonRepo(db)
    enrollmentRepo := postgres.NewEnrollmentRepo(db)

    h := &handler.Handler{
        Users:       userRepo,
        Courses:     courseRepo,
        Lessons:     lessonRepo,
        Enrollments: enrollmentRepo,
    }

    r := gin.Default()

    r.GET("/user/:id", h.User)
    r.POST("/user", h.UserCreate)
    r.PUT("/user/:id", h.UserUpdate)
    r.DELETE("/user/:id", h.UserDelete)

    r.GET("/course/:id", h.Course)
    r.POST("/course", h.CourseCreate)
    r.PUT("/course/:id", h.CourseUpdate)
    r.DELETE("/course/:id", h.CourseDelete)

    r.GET("/lesson/:id", h.Lesson)
    r.POST("/lesson", h.LessonCreate)
    r.PUT("/lesson/:id", h.LessonUpdate)
    r.DELETE("/lesson/:id", h.LessonDelete)

    r.GET("/enrollment/:id", h.EnrollmentId)
    r.POST("/enrollment", h.EnrollmentCreate)
    r.PUT("/enrollment/:id", h.EnrollmentUpdate)
    r.DELETE("/enrollment/:id", h.EnrollmentDelete)

    if err := r.Run(":8085"); err != nil {
        log.Fatal("Unable to start:", err)
    }
}
