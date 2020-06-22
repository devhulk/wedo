package main

import (
	"fmt"
	"log"
	"net/http"

	"wedo.io/api"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/wedo/create", api.CreateList)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
