package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
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

var dataTemp RequestData
var dataStorage SignUpData

func main(){

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/",handleHome)
	router.Post("/",handleHome)
	router.Get("/signup", handleSignup)
	router.Post("/signup", handleSignup)


	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	server.ListenAndServe()



}