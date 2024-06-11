package storage

import (
	"database/sql"
	"log"

	model "github.com/Go11Group/Dilfuza/Gin/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	dsn := "user=postgres password=1234 dbname=dilfuza sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func CreateUser(user *model.User) (*model.User, error) {
	query := `INSERT INTO users (id, name, last_name) VALUES($1, $2, $3) 
  		RETURNING id, name, last_name`

	var respUser model.User
	if err := db.QueryRow(query, user.Id, user.FirstName, user.LastName).Scan(
		&respUser.Id, &respUser.FirstName, &respUser.LastName); err != nil {
		return &model.User{}, err
	}

	return &respUser, nil
}

func UpdateUser(userId string, user *model.User) (*model.User, error) {
	query := `UPDATE users SET name = $1, last_name = $2 WHERE id = $3
	RETURNING id, name, last_name`

	var respUser model.User
	if err := db.QueryRow(query, user.FirstName, user.LastName, userId).Scan(
		&respUser.Id, &respUser.FirstName, &respUser.LastName); err != nil {
		return &model.User{}, err
	}

	return &respUser, nil
}

func DeleteUser(userId string) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, userId)
	return err
}

func CreateProblem(problem *model.Problem) (*model.Problem, error) {
	query := `INSERT INTO problems (id, title, description, difficulty) VALUES($1, $2, $3, $4) 
	RETURNING id, title, description, difficulty`

	var respProblem model.Problem
	if err := db.QueryRow(query, problem.Id, problem.Title, problem.Description, problem.Difficulty).Scan(
		&respProblem.Id, &respProblem.Title, &respProblem.Description, &respProblem.Difficulty); err != nil {
		return &model.Problem{}, err
	}

	return &respProblem, nil
}

func UpdateProblem(problemId string, problem *model.Problem) (*model.Problem, error) {
	query := `UPDATE problems SET title = $1, description = $2, difficulty = $3 WHERE id = $4
	RETURNING id, title, description, difficulty`

	var respProblem model.Problem
	if err := db.QueryRow(query, problem.Title, problem.Description, problem.Difficulty, problemId).Scan(
		&respProblem.Id, &respProblem.Title, &respProblem.Description, &respProblem.Difficulty); err != nil {
		return &model.Problem{}, err
	}

	return &respProblem, nil
}

func DeleteProblem(problemId string) error {
	query := `DELETE FROM problems WHERE id = $1`

	_, err := db.Exec(query, problemId)
	return err
}

