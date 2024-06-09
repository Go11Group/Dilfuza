package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Go11Group/Dilfuza/Gorilla_mux/models"
	"github.com/Go11Group/Dilfuza/Gorilla_mux/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/user/create", CreateUser).Methods("POST")
	r.HandleFunc("/user/update", UpdateUser).Methods("PUT")
	r.HandleFunc("/user/delete", DeleteUser).Methods("DELETE")

	// Problem routes
	r.HandleFunc("/problem/create", CreateProblem).Methods("POST")
	r.HandleFunc("/problem/update", UpdateProblem).Methods("PUT")
	r.HandleFunc("/problem/delete", DeleteProblem).Methods("DELETE")

	//r.HandleFunc("/solved_problem/create", CreateSolvedProblem).Methods("POST")

	log.Println("Server is running...")
	if err := http.ListenAndServe("localhost:8088", r); err != nil {
		log.Println("Error server is running!")
		return
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *models.User
	if err = json.Unmarshal(bodyByte, &user); err != nil {
		log.Println("error while unmarshalling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	user.Id = id

	respUser, err := storage.CreateUser(user)
	if err != nil {
		log.Println("error while creating user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *models.User
	if err = json.Unmarshal(bodyByte, &user); err != nil {
		log.Println("error while unmarshalling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user_id := r.URL.Query().Get("id")

	respUser, err := storage.UpdateUser(user_id, user)
	if err != nil {
		log.Println("error while updating user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user_id := r.URL.Query().Get("id")

	if err := storage.DeleteUser(user_id); err != nil {
		log.Println("error while deleting user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted User"))
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var problem *models.Problem
	if err = json.Unmarshal(bodyByte, &problem); err != nil {
		log.Println("Error while unmarshalling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	problem.Id = id

	respProblem, err := storage.CreateProblem(problem)
	if err != nil {
		log.Println("Error while creating problem", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	respBody, err := json.Marshal(respProblem)
	if err != nil {
		log.Println("Error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func UpdateProblem(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var problem *models.Problem
	if err = json.Unmarshal(bodyByte, &problem); err != nil {
		log.Println("Error while unmarshalling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	problemId := r.URL.Query().Get("id")

	respProblem, err := storage.UpdateProblem(problemId, problem)
	if err != nil {
		log.Println("Error while updating problem", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respBody, err := json.Marshal(respProblem)
	if err != nil {
		log.Println("Error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	problemId := r.URL.Query().Get("id")

	if err := storage.DeleteProblem(problemId); err != nil {
		log.Println("Error while deleting problem", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted Problem"))
}
