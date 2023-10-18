package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", LogMid(Home))
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("exec home handler")
}
func LogMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware is called")
		log.Println("doing middleware specific things")
		log.Println(r.Method)
		next(w, r)
		log.Println(r.URL)
	}
}
