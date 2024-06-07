package main

import (
	
	"net/http"
	"github.com/Go11Group/Dilfuza/lesson34/handler"
)

func main() {
	mux := http.ServeMux{}
	handler.Student(&mux)
	handler.Course(&mux)
	
	err := http.ListenAndServe(":8999", &mux)
	if err != nil {
		panic(err)
	}
}
