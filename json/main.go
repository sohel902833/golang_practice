package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSOn video")
	json := EnocodeJson()

	DecodeJson(json)
}

func EnocodeJson() []byte {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"Web-dev", "js"}},
		{"Mearn Bootcamp", 388, "LearnCodeOnline.in", "abc123", []string{"Web-dev", "js"}},
		{"Angular Bootcamp", 233, "LearnCodeOnline.in", "abc123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", finalJson)
	return finalJson
}

func DecodeJson(jsonDataFromWeb []byte) {
	fmt.Printf("%s\n", jsonDataFromWeb)
	checkValid := json.Valid(jsonDataFromWeb)
	var lcoCourse []course
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valild")
	}
}


func HelloWorld(){
	fmt.Println("Hello World!")
     name:="Md Sohab Hossain Sohel";
	 if name=="Md Sohab Hossain Sohel"{
		fmt.Println("Name Matched..")
	 }else if name=="Md Khorshed Alom Sujon"{
		fmt.Println("Name Doesn't Matched.")
	 }else if name=="Md Sohrab Alom Sujon"{
		fmt.Println("Name Doesn't Matched.")
	 }
}
