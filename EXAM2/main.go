package main

import (
	"log"

	"github.com/Go11Group/Dilfuza/EXAM2/api"
	"github.com/Go11Group/Dilfuza/EXAM2/api/handler"
	"github.com/Go11Group/Dilfuza/EXAM2/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Println("error while opening database", err)
	}
	defer db.Close()

	h := handler.NewHandler(
		postgres.NewUserRepo(db),
		postgres.NewCourseRepo(db),
		postgres.NewLessonRepo(db),
		postgres.NewEnrollmentRepo(db),
	)

	r := api.NewGin(h)
	err = r.Run(":8099")
	if err != nil {
		log.Println("error while starting server")
	}

	
}
