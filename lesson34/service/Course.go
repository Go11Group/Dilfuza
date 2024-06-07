package service

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/Go11Group/Dilfuza/lesson34/storage/postgres"

	"github.com/Go11Group/Dilfuza/lesson34/model"
	_ "github.com/gorilla/mux"
)

func CourseHandlerRepo() *postgres.CourseRepository {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	coursInfo := postgres.NewCourseRepository(db)
	return coursInfo
}
func CourseReadAllHandler(w http.ResponseWriter, r *http.Request) {
	coursInfo := CourseHandlerRepo()
	courses, err := coursInfo.ReadAllCourse()
	fmt.Println(courses)
	data, err := json.Marshal(courses)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusInternalServerError))))
	}
	_, err = w.Write(data)

}

func CourseDeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]

	coursInfo := CourseHandlerRepo()
	err := coursInfo.DeleteCourse(id)
	if err != nil {
		fmt.Println("salom", err)
		http.Error(w, "course is not  deleted", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("course is deleted"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {

	coursInfo := CourseHandlerRepo()
	courses := model.Course{}
	err := json.NewDecoder(r.Body).Decode(&courses)
	fmt.Println(courses)
	err = coursInfo.CreateUniversity(courses)
	if err != nil {
		fmt.Println("golang", err)
		_, err = w.Write([]byte("Is not create course"))
	} else {
		_, err = w.Write([]byte("Is  create course"))
	}
}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]
	coursInfo := CourseHandlerRepo()

	course := model.Course{}
	err := json.NewDecoder(r.Body).Decode(&course)
	fmt.Println(course)
	if err != nil {
		_, err = w.Write([]byte("Is not updated course"))
	}

	err = coursInfo.UpdateUniversity(id, course)
}
