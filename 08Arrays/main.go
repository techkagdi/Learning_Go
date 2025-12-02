package main

import "fmt"

func main() {
	fmt.Println("Lets learn an Array Using GO")

	var namelist [4]string
	namelist[0] = "Ankit"
	namelist[1] = "Sanket"
	namelist[3] = "Rohit"
	fmt.Println("Names are: ", namelist)
	fmt.Println("Lengths are  are: ", len(namelist))
}
