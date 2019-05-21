package main

import "github.com/kgcaHub/gocrud/db"
import "github.com/kgcaHub/gocrud/entity"

func main() {
	var person entity.Person
	person.Id = 123
	person.Name = "Klinsmann"

	db.Insert(person)
}
