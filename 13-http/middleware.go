package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", Mid(Mid1(HomePage)))
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page invoked")
	fmt.Fprintln(w, "this is my home")
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware invoked")
		next(w, r)
	}
}

func Mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 2 invoked")
		next(w, r)
	}
}
