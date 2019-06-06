package db

import (
	"database/sql"
	"fmt"

	//aun queda pendiente averiguar el guion
	_ "github.com/denisenkom/go-mssqldb"
)

var cxdb *sql.DB

//Open new MSSQL Connection
func NewConnection() {
	var errdb error
	cxdb, errdb = sql.Open("mssql", "server=.; user id=crudUser;password=crud2019*;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	fmt.Println("Se conectó")
}

//Close and dispose MSSQL Connection
func CloseConnection() {
	defer cxdb.Close()
	fmt.Println("Se cerró")
}
