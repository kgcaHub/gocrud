package db

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/kgcaHub/gocrud/entity"
)

//Insert a person
func Insert(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("NOR.Person_Insert ?,?", person.Id, person.Name)
	if err != nil {
		fmt.Println(" Insert error:", err.Error())
	}

	CloseConnection()
}
