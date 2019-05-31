package db

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/kgcaHub/gocrud/entity"
)

//Insert a person
func PersonCreate(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("NOR.Person_Insert ?,?", person.Id, person.Name)
	if err != nil {
		fmt.Println(" Insert error:", err.Error())
	}

	CloseConnection()
}

func PersonRead() {
	NewConnection()

	

	CloseConnection()
}

//Insert a person
func PersonUpdate(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("NOR.Person_Update ?,?", person.Id, person.Name)
	if err != nil {
		fmt.Println(" Update error:", err.Error())
	}

	CloseConnection()
}

func PersonDelete(person entity.Person) {
	NewConnection()

	_, err := cxdb.Exec("NOR.Person_Delete ?", person.Id)
	if err != nil {
		fmt.Println(" Delete error:", err.Error())
	}

	CloseConnection()
}