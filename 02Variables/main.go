package main

import "fmt"

const Pi float64 = 3.14

func main() {
	var username string = "Juzar"
	fmt.Println(username)
	fmt.Printf("Hello, %T!\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Booleans, %T!\n", isLoggedIn)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("Unsigned Integers, %T!\n", smallvalue)

	var website = "golong.org"
	fmt.Println(website)
	nuber := 123
	fmt.Println(nuber)

	fmt.Println(Pi)

}
