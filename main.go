package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	Amount    string `json:"amount"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Create new user
			user := User{
				FirstName: r.FormValue("fname"),
				LastName:  r.FormValue("lname"),
				Email:     r.FormValue("email"),
				Amount:    r.FormValue("amount"),
			}

			// Convert user to JSON
			jsonUser, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Send user data to REST API
			_, err = http.Post("http://localhost:8080/create", "application/json", bytes.NewBuffer(jsonUser))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return success response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("User created successfully!"))
		} else {
			// Render HTML form
			htmlBytes, err := ioutil.ReadFile("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			html := string(htmlBytes)
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(html))
		}
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
