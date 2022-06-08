package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome To Web Verb ")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/users"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count Is : ", byteCount)
	fmt.Println(responseString.String())
	fmt.Println(string(content))
}
