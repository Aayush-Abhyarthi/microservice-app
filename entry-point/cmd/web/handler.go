package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("templates/homepage.html", "templates/test.html"))

func handleHome(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "homepage.html", nil)
		if err != nil {
			http.Error(w, "Error loading the HTML template", http.StatusInternalServerError)
			fmt.Println("Error parsing HTML template:", err)
			return
		}
		return
	}

	if r.Method == http.MethodPost {

		var requestData RequestData
		requestData.UserID = r.FormValue("userId")
		requestData.Password = r.FormValue("password") 

		fmt.Print(requestData)

		sendingDataJson, err := json.Marshal(requestData)
		if err != nil {
			fmt.Println("An error occured while marshaling to json.")
		}

		url := "http://localhost:8000/"

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(sendingDataJson))
		if err != nil {
			fmt.Println("Error occured while doing post request")
		}

		defer resp.Body.Close()

		var respData ResponseData

		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			fmt.Println("Some error occured1")
		}

		fmt.Println(respData)

		if respData.IsValid {
			// Serve the success template
			fmt.Println("Inside the condition")
			err := templates.ExecuteTemplate(w, "test.html", nil)
			if err != nil {
				http.Error(w, "Error loading success page", http.StatusInternalServerError)
				fmt.Println("Error executing success template:", err)
			}
		} else {
			// Serve a failure message
			w.Write([]byte("User not authenticated"))
		}

	}
}
