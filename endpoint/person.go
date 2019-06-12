package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kgcaHub/gocrud/db"
	"github.com/kgcaHub/gocrud/entity"
)

func PersonCreate(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var person entity.Person

	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&person)

	if err != nil {
		fmt.Println("Decoding error:", err.Error())
	}
	fmt.Println(person)
	db.PersonCreate(person)
}

func PersonRead(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	persons := db.PersonRead()

	json.NewEncoder(w).Encode(persons)
}

func PersonUpdate(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var person entity.Person

	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&person)

	if err != nil {
		fmt.Println("Decoding error:", err.Error())
	}

	db.PersonUpdate(person)
}

func PersonDelete(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var person entity.Person

	jsonDecoder := json.NewDecoder(r.Body)
	err := jsonDecoder.Decode(&person)

	if err != nil {
		fmt.Println("Decoding error:", err.Error())
	}

	db.PersonDelete(person)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
