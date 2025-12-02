package main

import "fmt"

func main() {
	fmt.Println("Hello, World! Lets today Lear n Pointers in Golang")
	var oen int = 2
	var ptr *int = &oen
	fmt.Println("Value of oen is", oen)
	fmt.Println("Address of oen is ", &oen)
	fmt.Println("Value of the Pointer is ", ptr)

}
