package handler

import (
	"github.com/Go11Group/Dilfuza/EXAM2/postgres"
)
type Handler struct {
	UP  *postgres.UserRepo
	CP  *postgres.CourseRepo
	LP *postgres.LessonRepo
	EP  *postgres.EnrollmentRepo

}

func NewHandler(up *postgres.UserRepo, cp *postgres.CourseRepo, lp *postgres.LessonRepo, ep *postgres.EnrollmentRepo) *Handler {
	return &Handler{UP: up, CP: cp, LP: lp, EP: ep}
}
