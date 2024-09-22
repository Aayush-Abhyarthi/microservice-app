package main

import (
	"fmt"
	"net/http"
	

	"github.com/go-chi/chi/v5"
)

type RequestData struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type ResponseData struct {
	UserID string `json:"userID"`
	IsValid bool `json:"isValid"`
}

var mapVal map[string]bool


func main() {
	router := chi.NewRouter()
	router.Get("/", handleHome)
	router.Post("/", handleHome)
	router.Post("/login", handleLogin)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server started at http://localhost:3000")
	server.ListenAndServe()
}
