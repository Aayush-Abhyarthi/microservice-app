package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"golang.org/toolchain/src/math/rand"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("templates/homepage.html")
		if err != nil {
			http.Error(w, "Error loading the HTML template", http.StatusInternalServerError)
			fmt.Println("Error parsing HTML template:", err)
			return
		}
		temp.Execute(w, nil)
		return
	}


	if r.Method == http.MethodPost {
	
		var requestData RequestData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestData)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		fmt.Print(requestData.UserID)	
		
		sendingDataJson, err := json.Marshal(requestData)
		if(err!=nil){
			fmt.Println("An error occured while marshaling to json.")
		}

		url := "http://localhost:8000/"

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(sendingDataJson))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Response status:", resp.Status)
		
		rand.Seed(time.Now().UnixNano())
		randomValue := rand.Intn(100) 

	
		responseMessage := fmt.Sprintf("Hello, %s! Your random value is: %d", requestData.UserID, randomValue)

		
		responseData := ResponseData{Message: responseMessage}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	}
}