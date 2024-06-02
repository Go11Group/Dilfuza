package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string
	Age     int
	Courses []Course `gorm:"many2many:student_courses;"`
}

type Course struct {
	gorm.Model
	CourseName string
	Price      int
	Students   []Student `gorm:"many2many:student_courses;"`
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=dilfuza sslmode=disable TimeZone=Asia/Tashkent"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&Student{}, &Course{}); err != nil {
		panic(err)
	}

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. View all students")
		fmt.Println("2. Insert a new student and course")
		fmt.Println("3. Update a student")
		fmt.Println("4. Delete a student")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			students, err := getAllUsers(db)
			if err != nil {
				fmt.Println("Error fetching students:", err)
			} else {
				fmt.Println("\n--- All Students ---")
				for _, student := range students {
					fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
					if len(student.Courses) > 0 {
						fmt.Println("  Courses:")
						for _, course := range student.Courses {
							fmt.Printf("    - %s\n", course.CourseName)
						}
					}
				}
			}

		case 2:
			insertFunc(db)

		case 3:
			var studentID int
			fmt.Print("Enter student ID to update: ")
			fmt.Scan(&studentID)
			err := updateStudent(db, studentID)
			if err != nil {
				fmt.Println("Error updating student:", err)
			}

		case 4:
			var studentID int
			fmt.Print("Enter student ID to delete: ")
			fmt.Scan(&studentID)
			err := deleteStudent(db, studentID)
			if err != nil {
				fmt.Println("Error deleting student:", err)
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}

func insertFunc(db *gorm.DB) {
	var studentName string
	var studentAge int
	fmt.Print("Enter student's name: ")
	fmt.Scan(&studentName)
	fmt.Print("Enter student's age: ")
	fmt.Scan(&studentAge)

	student := Student{
		Name: studentName,
		Age:  studentAge,
	}
	db.Create(&student)

	var courseName string
	var coursePrice int
	fmt.Print("Enter course name: ")
	fmt.Scan(&courseName)
	fmt.Print("Enter price: ")
	fmt.Scan(&coursePrice)

	course := Course{
		CourseName: courseName,
		Price:      coursePrice,
	}
	db.Create(&course)

	db.Model(&student).Association("Courses").Append(&course)

	fmt.Println("Inserted successfully")
}

func getAllUsers(db *gorm.DB) ([]Student, error) {
	var students []Student
	err := db.Model(&Student{}).Preload("Courses").Find(&students).Error
	return students, err
}

func getOneUser(db *gorm.DB, studentID int) (Student, error) {
	var student Student
	err := db.Preload("Courses").First(&student, studentID).Error
	return student, err
}

func updateStudent(db *gorm.DB, studentID int) error {
	student, err := getOneUser(db, studentID)
	if err != nil {
		return err
	}

	var newName string
	var newAge int
	fmt.Print("Enter new name: ")
	fmt.Scan(&newName)
	fmt.Print("Enter new age: ")
	fmt.Scan(&newAge)

	student.Name = newName
	student.Age = newAge

	err = db.Save(&student).Error
	if err != nil {
		return err
	}

	fmt.Println("Update successful")
	return nil
}

func deleteStudent(db *gorm.DB, studentID int) error {
	student, err := getOneUser(db, studentID)
	if err != nil {
		return err
	}

	err = db.Delete(&student).Error
	if err != nil {
		return err
	}

	fmt.Println("Delete successful")
	return nil
}