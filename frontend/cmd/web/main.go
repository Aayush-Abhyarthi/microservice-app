package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type RequestData struct {
    Data string `json:"data"`
}

func handleHome(w http.ResponseWriter, r *http.Request){

	if(r.Method == http.MethodPost){
		
		// Parse the JSON body
		var requestData RequestData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
		}


	}
	temp,err := template.ParseFiles("templates/homepage.html")
	if(err!=nil){
		fmt.Println("There was an error in the parsing of html document")
	}

	temp.Execute(w,nil)

	
}

func main(){

	router := chi.NewRouter()
	router.Get("/", handleHome)
	router.Post("/", handleHome)

	server := &http.Server{
		Addr : ":3000",
		Handler : router,
	}

	server.ListenAndServe()
}