package api

import (
	"log"
	"net/http"
)

func main() {

}

func handleRequests() {

	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":5007", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
