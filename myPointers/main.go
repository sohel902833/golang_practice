package main

import "fmt"

func main() {
	fmt.Print("Welcome to Pointer Class.")

	myNumber := 10
	//var number *int = &myNumber

	fmt.Println("Before Change", myNumber)
	changeNumber(&myNumber)
	fmt.Println("After Changed ", myNumber)
}

func changeNumber(num *int) {
	*num = *num * 20
}
