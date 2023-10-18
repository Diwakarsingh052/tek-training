// import necessary packages
package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Creating a custom type for context key
type reqKey int

// A constant for request id key
const RequestIDKey reqKey = 123

// Main execution function
func main() {
	// Assign endpoint "/home" to be handled by the layered handlers
	http.HandleFunc("/home", RequestIdMid(LoggingMid(homePage)))
	// Start running the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Function to handle requests at homePage
func homePage(w http.ResponseWriter, r *http.Request) {
	// Print logs for each request received
	log.Println("In home Page handler")
	// Respond to the client request
	fmt.Fprintln(w, "this is my home page")
}

// Middleware function to generate and assign unique request ID

func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate UUID for the request
		uuid := uuid.NewString()
		// Assign the UUID to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, uuid)
		// Call the next function in the middleware chain with the updated context
		next(w, r.WithContext(ctx))
	}
}

// Middleware function to log request details

func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the request ID from the context
		ctx := r.Context()
		reqId, ok := ctx.Value(RequestIDKey).(string)
		// If request ID not found, assign default "Unknown"
		if !ok {
			reqId = "Unknown"
		}
		// Log the details of the request
		log.Printf("%s : started   : %s %s ",
			reqId,
			r.Method, r.URL.Path)
		// Ensure the completion log line is printed even in case of panics
		defer log.Println("completed")
		// Call the next function in the middleware chain
		next(w, r)
	}
}
