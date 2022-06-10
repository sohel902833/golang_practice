package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://www.youtube.com/"

func main() {
	fmt.Println("Welcome to LCO web request")
	performPostJsonRequest()
}

func performGetRequest() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T\n", response)
	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}

func performPostJsonRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`
		{
			"name":"Md Sohrab Hossain Sohel",
			"roll":902833,
			"platform":"SohelCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
func performPostFormRequest() {

}
