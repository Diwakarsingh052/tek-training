// Package `handlers` contains all http handlers associated with different endpoints
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proj5/models" // import the models package where all db access functions reside
	"strconv"
)

// GetUser is the HTTP handler for fetching a user record.
// The function expects a user_id parameter from the URL.
// It sends back a user record if agreement is found or error messages if user not found or an internal error occurred.
func GetUser(w http.ResponseWriter, r *http.Request) {

	// Fetch the user_id from the url parameters
	userIdString := r.URL.Query().Get("user_id")

	// Convert the user_id string fetched from url to uint64 data type
	userId, err := strconv.ParseUint(userIdString, 10, 64)

	// If the conversion fails, respond with an error
	if err != nil {
		log.Println("Error: ", err)
		// Prepare an error message in a map[string]string
		errorInConversion := map[string]string{"msg": "not a valid number"}

		// Convert the error message into JSON format
		jsonData, err := json.Marshal(errorInConversion)

		// If JSON conversion fails, respond with internal server error
		if err != nil {
			log.Println("Error while converting error to json", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		// Respond with a bad request response code and the error message
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}

	// Fetch user data from the database using the user id
	uData, err := models.FetchUser(userId)

	// If user fetch operation fails, respond with an error
	if err != nil {
		fetchError := map[string]string{"msg": "user not found"}
		errData, err := json.Marshal(fetchError)
		if err != nil {
			log.Println("Error while parsing fetchuser error conversion: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		// Respond with an internal server error code and the error message
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errData)
		return
	}

	// Convert the user data into JSON format
	userData, err := json.Marshal(uData)

	// If user data JSON conversion fails, respond with internal server error
	if err != nil {
		log.Println("Error while converting user data to json", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Respond with user data
	w.Write(userData)
}
