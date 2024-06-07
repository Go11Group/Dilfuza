
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var fruits = map[string]string{"1": "apple", "54": "banana", "23": "lemon"}

func main() {
	http.HandleFunc("/fruit/", juice)
	http.HandleFunc("/concatenate-full-name", concatenation)
	http.HandleFunc("/book", book)
	http.HandleFunc("/car", car)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/student", student)
	http.HandleFunc("/users", users)

	err := http.ListenAndServe(":8117", nil)
	if err != nil {
		panic(err)
	}
}

func juice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Malumot topaolmadim ", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.URL.Path)
	w.Write([]byte(fruits[strings.TrimPrefix(r.URL.Path, "/fruit/")]))
}

func concatenation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Malumot topaolmadim ", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(r.URL.Query().Get("name") + " " + r.URL.Query().Get("lastname")))
}

func book(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Malumot topaolmadim", http.StatusMethodNotAllowed)
		return
	}
	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding, err: %s", err.Error())))
		return
	}

	fmt.Println(b)
	w.Write([]byte("SUCCESS"))
}

type Book struct {
	Name      string
	Author    string
	Publisher string
}

func car(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Ma'lumot topaolmadim ", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("SUCCESS"))
}
type Car struct{
	name string
	year string
}

func cat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Ma'lumot topaolmadim", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("Cat API response"))
}

func student(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Ma'lumot topaolmadim ", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("Student API response"))
}

func users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Ma'lumot topaolmadim", http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("Users API response"))
}
