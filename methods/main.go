package main

import "fmt"

func main() {
	fmt.Println("Welcome To Functions.")
	sohel := User{"Sohel", "sohel@gmail.com", false, 20, 45}
	sohel.GetStatus()
	sohel.NewMail()
	fmt.Println("Prev Email: ", sohel.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}
func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("new Email ", u.Email)

}
