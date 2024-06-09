package storage

import (
	"database/sql"

	"github.com/Go11Group/Dilfuza/Gorilla_mux/models"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	dsn := "user=postgres password=1234 dbname=dilfuzadb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return &models.User{}, err
	}

	defer db.Close()

	query := `INSERT INTO users (id, name, last_name) VALUES($1, $2, $3) 
  	RETURNING id, name, last_name`

	var respUser models.User
	if err = db.QueryRow(query,
		user.Id,
		user.FirstName,
		user.LastName).Scan(&respUser.Id, &respUser.FirstName, &respUser.LastName); err != nil {
		return &models.User{}, err
	}

	return &respUser, nil
}

func UpdateUser(userId string, user *models.User) (*models.User, error) {
	db, err := connect()
	if err != nil {
		return &models.User{}, err
	}
	defer db.Close()

	query := `
  	UPDATE 
    	users 
  	SET 
    	name = $1, 
    	last_name = $2
  	WHERE 
    	id = $3
  	RETURNING 
    	id, 
    	name, 
    	last_name`

	var respUser models.User
	if err := db.QueryRow(query, user.FirstName, user.LastName, userId).Scan(
		&respUser.Id,
		&respUser.FirstName,
		&respUser.LastName,
	); err != nil {
		return &models.User{}, err
	}

	return &respUser, nil
}

func DeleteUser(userId string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM users WHERE id = $1`

	_, err = db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func CreateProblem(problem *models.Problem) (*models.Problem, error) {
	db, err := connect()
	if err != nil {
		return &models.Problem{}, err
	}
	defer db.Close()

	query := `INSERT INTO problems (id, title, description, difficulty) VALUES($1, $2, $3, $4) 
	RETURNING id, title, description, difficulty`

	var respProblem models.Problem
	if err = db.QueryRow(query, problem.Id, problem.Title, problem.Description, problem.Difficulty).Scan(
		&respProblem.Id, &respProblem.Title, &respProblem.Description, &respProblem.Difficulty); err != nil {
		return &models.Problem{}, err
	}

	return &respProblem, nil
}

func UpdateProblem(problemId string, problem *models.Problem) (*models.Problem, error) {
	db, err := connect()
	if err != nil {
		return &models.Problem{}, err
	}
	defer db.Close()

	query := `UPDATE problems SET title = $1, description = $2, difficulty = $3 WHERE id = $4
	RETURNING id, title, description, difficulty`

	var respProblem models.Problem
	if err = db.QueryRow(query, problem.Title, problem.Description, problem.Difficulty, problemId).Scan(
		&respProblem.Id, &respProblem.Title, &respProblem.Description, &respProblem.Difficulty); err != nil {
		return &models.Problem{}, err
	}

	return &respProblem, nil
}

func DeleteProblem(problemId string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM problems WHERE id = $1`

	_, err = db.Exec(query, problemId)
	if err != nil {
		return err
	}
	return nil
}
/*
func CreateSolvedProblem(solvedProblem *models.SolvedProblem) (*models.SolvedProblem, error) {
	db, err := connect()
	if err != nil {
		return &models.SolvedProblem{}, err
	}
	defer db.Close()

	query := `INSERT INTO solved_problems (user_id, problem_id, solved_at) VALUES($1, $2, $3) 
	RETURNING user_id, problem_id, solved_at`

	var respSolvedProblem models.SolvedProblem
	if err = db.QueryRow(query, solvedProblem.UserId, solvedProblem.ProblemId, solvedProblem.SolvedAt).Scan(
		&respSolvedProblem.UserId, &respSolvedProblem.ProblemId, &respSolvedProblem.SolvedAt); err != nil {
		return &models.SolvedProblem{}, err
	}

	return &respSolvedProblem, nil
}
*/