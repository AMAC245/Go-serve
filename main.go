package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, port())
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sub")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", subHandler)
	http.ListenAndServe(port(), nil)
}
