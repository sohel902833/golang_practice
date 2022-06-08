package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gasdfasdfweridf"

func main() {
	fmt.Println("Welcome to handling URLs in golan")
	fmt.Println(myurl)
	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qParams := result.Query()
	fmt.Println(qParams["paymentid"])

	for _, val := range qParams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sohrab",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
