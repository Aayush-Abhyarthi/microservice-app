package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handleGetInfo)
	router.Post("/", handleGetInfo)

	server := &http.Server{
		Addr: "8070",
		Handler: router,
	}

	server.ListenAndServe()

}