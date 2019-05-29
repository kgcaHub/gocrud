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

	db.PersonCreate(person)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
