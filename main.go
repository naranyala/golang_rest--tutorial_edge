package main

import (
	"encoding/json"
	"fmt"
	"gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title ",
			Desc:    "Test Description",
			Content: "Hello World",
		},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	//log.Fatal(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequest()
}
