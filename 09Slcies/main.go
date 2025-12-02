package main

import "fmt"

func main() {
	fmt.Println("Hello, World! Lets today Learn Slices in Golang")
	var fru = []int{23, 45, 67, 89}

	fmt.Println("Numbers are: ", fru)
	fru = append(fru, 95)
	fmt.Println("Appending the Slice using Append Method", fru)
	fru = append(fru[1:], 95)
	fmt.Println("Slicing the SLice is wiht 1: syntax ", fru)

	fru = append(fru[1:3])
	fmt.Println("Slicing the SLice is wiht 1:3 syntax ", fru)
}
