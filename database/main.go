package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleGetInfo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is the value in the main file"))
}

func main(){

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handleGetInfo)
	router.Post("/", handleGetInfo)

	server := &http.Server{
		Addr: "8080",
		Handler: router,
	}

	server.ListenAndServe()

}