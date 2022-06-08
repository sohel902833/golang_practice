package main

import "fmt"

func main() {
	fmt.Println("Welcome To Maps")

	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)
	fmt.Println("", languages["Js"])

	//loops
	for key, value := range languages {
		fmt.Println(key, value)
	}

}
