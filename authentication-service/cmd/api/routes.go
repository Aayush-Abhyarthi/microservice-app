package main

import (
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

		fmt.Println("The response object to be sent to backend")
		fmt.Println(authenticationPayload)

		//

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		
			// Encode and send JSON response
			if err := json.NewEncoder(w).Encode(authenticationPayload) 
			err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		//



		//end sending

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received successfully"))

	}



}
