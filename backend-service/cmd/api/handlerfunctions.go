package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var reqD RequestData
		if err := json.NewDecoder(r.Body).Decode(&reqD); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		fmt.Println(reqD)

		// send reqD to authorization-service

		jsonData, err := json.Marshal(reqD)
		if err != nil {
			fmt.Println("Error occured")
		}

		url := "http://localhost:8080"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}

		//read output from the authorization service
		var objectResponse ResponseData
		err = json.NewDecoder(resp.Body).Decode(&objectResponse)
		if err != nil {
			fmt.Println("An error occured")
		}

		fmt.Println("Received authpayload")
		fmt.Println(objectResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(objectResponse)

	}

}
