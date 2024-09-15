package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func handleHome(w http.ResponseWriter, r *http.Request){
	temp,err := template.ParseFiles("templates/homepage.html")
	if(err!=nil){
		fmt.Println("There was an error in the parsing of html document")
	}

	temp.Execute(w,nil)

	
}

func main(){

	router := chi.NewRouter()
	router.Get("/", handleHome)
	//router.POST("/", handleHome)

	server := &http.Server{
		Addr : ":3000",
		Handler : router,
	}

	server.ListenAndServe()
}