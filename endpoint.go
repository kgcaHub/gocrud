package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kgcaHub/gocrud/endpoint"
)

func main() {
	handleRequests()
	fmt.Scanln()
}

func handleRequests() {

	http.Handle("/", http.FileServer(http.Dir("./ui")))
	http.HandleFunc("/person/create", endpoint.PersonCreate)
	log.Fatal(http.ListenAndServe(":5007", nil))
}
