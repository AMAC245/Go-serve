package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//typed routes/api
const (
	base         string = "/"
	sub          string = "/sub"
	endpoint     string = "https://api.dribbble.com/v1/shots"
	test_rebound string = "/3850371"
	access       string = "?access_token=859f432cad2e206603ed29d9c5728f0c14673e2ea14c453b07f3b05a36cd2fea"
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
	resp, err := http.Get(endpoint + access)
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
	http.HandleFunc(base, indexHandler)
	http.HandleFunc(sub, subHandler)
	http.ListenAndServe(port(), nil)
}
