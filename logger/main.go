package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func handleLog(w http.ResponseWriter, r *http.Request){

	

}

func main(){

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/", handleLog)

	server := &http.Server{
		Addr: "4040",
		Handler : router,
	}

	server.ListenAndServe()

}