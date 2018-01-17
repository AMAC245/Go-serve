package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//typed routes/api
const (
	BASE         string = "/"
	SUB          string = "/sub"
	ENDPOINT     string = "https://api.dribbble.com/v1/shots"
	TEST_REBOUND string = "/3850371"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "base")
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

func fetchData() {
	token := os.Getenv("TOKEN")
	resp, err := http.Get(ENDPOINT + token)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	data := string(bytes)
	fmt.Println(data)

	resp.Body.Close()
}

func main() {
	fetchData()

	http.HandleFunc(BASE, indexHandler)
	http.HandleFunc(SUB, subHandler)
	http.ListenAndServe(port(), nil)
}
