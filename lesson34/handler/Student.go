package handler

import (
	"github.com/Go11Group/Dilfuza/lesson34/service"
	"net/http"
)

func Student(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/student/all", service.StudentReadAllHandler)
	mux.HandleFunc("POST /api/student/create", service.CreateStudentHandler)
	mux.HandleFunc("PUT /api/student/update", service.UpdateStudentHandler)
	mux.HandleFunc("DELETE /api/student/delete", service.DeleteStudentHandler)

}
