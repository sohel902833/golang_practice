package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome To Files in goloang")

	content := "This needs to go in a file -Md Sohrab Hossain Sohel"

	file, err := os.Create("./server.txt")

	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length Is: ", length)
	defer file.Close()
	readFile("./server.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Byte Data: ", databyte)
	fmt.Println("Data Inside the file is \n", databyte)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
