package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func handleHome(w http.ResponseWriter, r *http.Request){

	if(r.Method == http.MethodPost){

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var dataTemp RequestData

		err = json.Unmarshal(body, &dataTemp)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received object: %+v\n", dataTemp)


		//Sending json to authentication-service

		url := "http://localhost:8080/"
		jsonDataTemp,err := json.Marshal(dataTemp)
		if(err!=nil){
			fmt.Println("Error occured during marshaling")
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonDataTemp))
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

		//end sending

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received successfully"))

	}

}

func handleAuth(w http.ResponseWriter, r *http.Request){

	if(r.Method == http.MethodPost){
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

		fmt.Printf("Received object from auth: %+v\n", dataTemp)

		//
			url := "http://localhost:3000/login"
			jsonDataTemp,err := json.Marshal(dataTemp)
			if(err!=nil){
				fmt.Println("Error occured during marshaling")
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonDataTemp))
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
		//

	}

}