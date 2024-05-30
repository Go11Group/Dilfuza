package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

type Author struct {
    Id   int
    Name string
    Age  int
}

type Book struct {
    Id   int
    Name string
}

type Books_Authors struct {
    Id        int
    Author_Id Author
    Book_Id   Book
}

func main() {
  connStr := "user=postgres password=1234 dbname=postgres sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  addData(db)

  fetchData(db)

 displayData(db)
}

func addData(db *sql.DB) {
  _, err := db.Exec(`
    INSERT INTO Author (Id,Name, Age) VALUES
    (1,'Dilfuza', 35),
    (2,'Samina', 15),
    (3,'Odil', 20)
  `)
  if err != nil {
    panic(err)
  }

  _, err = db.Exec(`
    INSERT INTO Book (Id, Name) VALUES
    (1, 'book1'),
	(2,'book2'),
	(3, 'book3')
  `)
  if err != nil {
    panic(err)
  }

  _, err = db.Exec(`
    INSERT INTO Books_Authors (Id, Author_id,Book_id ) VALUES
    (1, 1, 3),
    (1, 2, 2),
    (1, 3, 1)
  `)
  if err != nil {
    panic(err)
  }
}

func fetchData(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM Author")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("Authors:")
    for rows.Next() {
        var author Author
        if err := rows.Scan(&author.Id, &author.Name, &author.Age); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("Id: %d, Name: %s, Age: %d\n", author.Id, author.Name, author.Age)
    }

    rows, err = db.Query("SELECT id, name FROM Book")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("\nBooks:")
    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.Id, &book.Name); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("Id: %d, Name: %s\n", book.Id, book.Name)
    }

    rows, err = db.Query("SELECT id, author_id, book_id FROM Books_Authors")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("\nBooks_Authors:")
    for rows.Next() {
        var books_authors Books_Authors
        if err := rows.Scan(&books_authors.Id, &books_authors.Author_Id, &books_authors.Book_Id); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("Id: %d, Author_Id: %d, Book_Id: %d\n", books_authors.Id, books_authors.Author_Id,books_authors.Book_Id)
    }
	
}

func displayData(db *sql.DB) {
    fetchData(db)
}
