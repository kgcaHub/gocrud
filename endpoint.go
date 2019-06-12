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
	http.HandleFunc("/", moduloServe)
	http.HandleFunc("/person/create", endpoint.PersonCreate)
	http.HandleFunc("/person/read", endpoint.PersonRead)
	http.HandleFunc("/person/update", endpoint.PersonUpdate)
	http.HandleFunc("/person/delete", endpoint.PersonDelete)
	log.Fatal(http.ListenAndServe(":5007", nil))
}

func moduloServe(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Server resource in")
	var _path = req.URL.Path
	fmt.Println(_path)

	if _path == "/" {
		fmt.Println("calling index")
		http.ServeFile(w, req, "ui/index.html")
	}

	if _path == "/main.js" {
		w.Header().Add("Content-Type", "text/javascript")
		http.ServeFile(w, req, "ui/main.js")
	}

	if _path == "/components/person.js" {
		w.Header().Add("Content-Type", "text/javascript")
		http.ServeFile(w, req, "ui/components/person.js")
	}

	if _path == "/components/person.html" {
		fmt.Println("calling person component")
		http.ServeFile(w, req, "ui/components/person.html")
	}
}
