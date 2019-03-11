package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	// "log"
	// "github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id 		int		`json:"Id"`
	Title 	string	`json:"Title"`
	Desc 	string	`json:"desc"`
	Content string 	`json:"content"`
}

var Article []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){

}

func main()  {
	
}