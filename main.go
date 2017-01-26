package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/ping", Pong)

	log.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to gilak</h1>")
}
