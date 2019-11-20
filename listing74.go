package main

import (
	"D1/ex1/entities"
	"fmt"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}
	a.Name = "bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
