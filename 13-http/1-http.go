package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Id      int
	Hobbies []string
}

func main() {

	http.HandleFunc("/home", home)
	//start your server
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {

	// Define a user with id and hobbies
	u := user{
		Id:      101,
		Hobbies: []string{"cricket", "football"},
	}

	// Marshal the user into JSON
	jsonData, err := json.Marshal(u)

	// If there was an error while marshaling
	if err != nil {
		// Log the error
		log.Println(err)

		// Respond with an internal server error status
		w.WriteHeader(http.StatusInternalServerError)

		// Write the internal server status text to the response body
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))

		// Stop the function
		return
	}

	// If there was no error, respond with 'OK' status
	w.WriteHeader(http.StatusOK)

	// Write the JSON data to the response body
	w.Write(jsonData)
}
