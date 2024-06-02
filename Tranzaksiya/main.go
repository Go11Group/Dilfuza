package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Book struct {
	Id    int
	Name  string
	Count int
}

func main() {
	connection := "user=postgres password=1234 dbname=dilfuza sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	CreateUserWithBook(db)
}

func UpdateBook(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	var Id, newCount int
	var newName string
	fmt.Print("Enter user id to update : ")
	fmt.Scan(&Id)
	fmt.Print("Enter new book name: ")
	fmt.Scan(&newName)
	fmt.Print("Enter new book count: ")
	fmt.Scan(&newCount)

	_, err = tx.Exec("UPDATE books SET name = $1, count = $2 WHERE id = $3", newName, newCount, Id)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	fmt.Print("Updated succesfully!!!")
}

func UpdateUser(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	var newName string
	var id, newAge int
	fmt.Print("Enter user id to update: ")
	fmt.Scan(&id)
	fmt.Print("Enter new name: ")
	fmt.Scan(&newName)
	fmt.Print("Enter new age: ")
	fmt.Scan(&newAge)
	_, err = tx.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", newName, newAge, id)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	fmt.Println("User updated succesfully!!!")
}

func CreateUserWithBook(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	var username, bookname string
	var userage, bookcount, respUserId int
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your age: ")
	fmt.Scan(&userage)
	fmt.Print("Enter book`s name: ")
	fmt.Scan(&bookname)
	fmt.Print("How many books: ")
	fmt.Scan(&bookcount)

	if err = tx.QueryRow("INSERT INTO users(name, age) VALUES ($1, $2) RETURNING id", username, userage).Scan(&respUserId); err != nil {
		tx.Rollback()
		panic(err)
	}

	_, err = tx.Exec("INSERT INTO books(name, count, user_id) VALUES ($1, $2, $3)", bookname, bookcount, respUserId)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	fmt.Print("User created succesfully!!!")

}