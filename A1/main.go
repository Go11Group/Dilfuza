package main

import (
	"A1/package2"
	"A1/packagee"
	"fmt"
)

func main() {
	person := packagee.Person{
		Name:    "Dilfuza",
		Surname: "Dilfuza",
		Age:     18,
	}

	fmt.Println(person.PrintName())
	fmt.Println(person.PrintSurname())

	package2.PrintHello("Dilfuza")

}
