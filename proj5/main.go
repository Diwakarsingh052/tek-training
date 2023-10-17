package main

import (
	"net/http"
	"proj5/handlers"
)

func main() {

	http.HandleFunc("/user", handlers.GetUser)
	panic(http.ListenAndServe(":8080", nil))

}
