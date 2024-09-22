package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"
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

		fmt.Println("Authenticating.... Please wait!!!!")
		time.Sleep(2 * time.Second)

		if(mapVal[requestData.UserID]==true){

			fmt.Println("User authenticated")

		}else{
			fmt.Println("User not authenticated")
		}

	}
}

func handleLogin(w http.ResponseWriter, r *http.Request){

	if(r.Method == http.MethodPost){

		fmt.Println("In the handleLogin function")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var dataTemp ResponseData

		err = json.Unmarshal(body, &dataTemp)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received object from backend: %+v\n", dataTemp)
		if mapVal == nil {
			mapVal = make(map[string]bool)  // Initialize the map
		}
		mapVal[dataTemp.UserID]=dataTemp.IsValid
		fmt.Println("Successfully executed the handleLogin function")
	}
}