package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	user = "postgres"
	dbname = "dilfuza"
	password = 1234
	port = 5432
)

func ConnectDB() (*sql.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s user=%s dbname=%s password=%d port=%d sslmode=disable", host, user, dbname, password, port)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}