package main

import "net/http"

func handleGetInfo(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is the output"))

}
