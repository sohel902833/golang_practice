package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

type ArticleResponse struct {
	articles []Article `json:"Article"`
	message  string    `json:"Message"`
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hellow World"},
	}

	// response := ArticleResponse{
	// 	articles: articles,
	// 	message:  "This Message For Article Response",
	// }

	fmt.Println("Endpoint Hits: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Endpoint Hit")
}
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func main() {
	handleRequests()
}
