package main

import "fmt"

func main() {

	fmt.Println("Lets Start with the Struct")
	student1 := Student{"JK", 33, 150}

	fmt.Println("Student Details are: ", student1)

	fmt.Printf("Student Details using Printf is %v\n", student1)
}

type Student struct {
	Name   string
	RollNo int
	Marks  int
}
