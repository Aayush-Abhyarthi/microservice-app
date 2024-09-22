package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-openapi/validate"
)

func handleHome(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err = json.Unmarshal(body, &dataTemp)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received object: %+v\n", dataTemp)

		//Sending json to authentication-service

		val := validateCreds(dataTemp)

		authenticationPayload := ResponseData{UserID: dataTemp.UserID ,IsValid: val }

		//

			url := "http://localhost:8000/authres"
			jsonDataTemp,err := json.Marshal(authenticationPayload)
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



		//end sending

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received successfully"))

	}



}
