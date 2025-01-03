package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestData struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type ResponseData struct {
	UserID string `json:"userID"`
	IsValid bool `json:"isValid"`
}

type SignUpData struct {
	UserName string `json:"userId"`
	Password string `json:"password"`
	GUserName string `json:"githubusername"`
}

var mapVal map[string]bool

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", handleHome)
	router.Post("/", handleHome)
	router.Get("/signup", handleSignup)
	router.Post("/signup", handleSignup)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}
