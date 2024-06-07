package handler

import (
	"github.com/Go11Group/Dilfuza/lesson34/service"
	"net/http"
)


func Course(mux *http.ServeMux) {
	mux.HandleFunc("GET/api/course/all", service.CourseReadAllHandler)
	mux.HandleFunc("POST /api/course/create", service.CreateCourseHandler)
	mux.HandleFunc("PUT /api/course/update", service.UpdateCourseHandler)
	mux.HandleFunc("DELETE /api/course/delete", service.CourseDeleteHandler)

}