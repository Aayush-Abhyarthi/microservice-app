package main

import (
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

		fmt.Printf("Received JSON: %+v\n", dataTemp)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JSON received successfully"))

	}

}