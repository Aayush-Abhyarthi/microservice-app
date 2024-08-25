package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		render(w, "test.page.gohtml")
	})

	fmt.Println("Server started on server 80");
	err := http.ListenAndServe(":80", nil)
	if(err!=nil){
		log.Panic(err)
	}
}