package main

import "fmt"

const LoginToken string = "asdfasdfsdfksdkfjskfjjd" //public variable

func main() {
	var username string = "Md Sohrab Hossain"
	fmt.Print(username)
	fmt.Printf(" Variable is of type : %T\n", username)

	var isLogggedIn bool = false
	fmt.Print(username)
	fmt.Printf(" Variable is of type : %T\n", isLogggedIn)

	var smallVal uint8 = 255
	fmt.Print(smallVal)
	fmt.Printf(" Variable is of type : %T\n", smallVal)

	var smallFloat float64 = 255.2342343
	fmt.Print(smallFloat)
	fmt.Printf(" Variable is of type : %T\n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Print(anotherVariable)
	fmt.Printf(" Variable is of type : %T\n", anotherVariable)

	//implicit type
	website := "learncodeonline.in"
	fmt.Print(website)
	fmt.Printf(" Variable is of type : %T\n", website)

	fmt.Print("Logged in token", LoginToken)
}
