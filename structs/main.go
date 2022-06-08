package main

import "fmt"

func main() {
	fmt.Println("Structs in golan")

	sohel := User{"Sohel", "sohel@gmail.com", true, 20}
	fmt.Println("Name", sohel.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
