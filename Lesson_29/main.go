
package main

import (
	"time"
	"github.com/k0kubun/pp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	//IsEmployee bool
	Birthday time.Time
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres  sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	users := User{ FirstName: "jin",LastName: "Dooooo",Email: "jindooooo@gmail.com",Password: "1234", Age: 19,Field: "Toshkent",Gender: "alo", Birthday: time.Now()}

	result := db.Create(&users)

	pp.Println(result.Error)
	
}
