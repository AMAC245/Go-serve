package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
}

func main() {
	//router instance
	route := mux.NewRouter()

	route.HandleFunc("/movies", getMovies).Method("GET")
	route.HandleFunc("/movies/{id}", getMovie).Method("GET")
	route.HandleFunc("/movies", addMovies).Method("GET")
	route.HandleFunc("/movies/{id}", updateMovie).Method("GET")
	route.HandleFunc("/movies/{id}", deleteMovies).Method("GET")

	log.Fatal(http.ListenAndServe(":5000", route))
}
