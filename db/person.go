package db

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/kgcaHub/gocrud/entity"
)

//Insert a person
func PersonCreate(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("CRUD.Person_Create ?,?", person.Id, person.Name)
	if err != nil {
		fmt.Println(" Insert error:", err.Error())
	}

	CloseConnection()
}

func PersonRead() []entity.Person {

	resultado := make([]entity.Person, 0)

	NewConnection()

	rows, err := cxdb.Query("CRUD.Person_Read")

	if err != nil {
		fmt.Println(" Read error:", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var person entity.Person
		if err := rows.Scan(&person.Id, &person.Name); err != nil {
			fmt.Println(err)
		}
		resultado = append(resultado, person)
	}

	CloseConnection()

	return resultado
}

//Update a person
func PersonUpdate(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("CRUD.Person_Update ?,?", person.Id, person.Name)
	if err != nil {
		fmt.Println(" Update error:", err.Error())
	}

	CloseConnection()
}

//Delete a person
func PersonDelete(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("CRUD.Person_Delete ?", person.Id)
	if err != nil {
		fmt.Println(" Delete error:", err.Error())
	}

	CloseConnection()
}
