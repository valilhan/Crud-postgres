package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	var PersonID, LastName, FirstName, Address, City string
	n, err := fmt.Scanln(&PersonID, &FirstName, &LastName, &Address, &City)
	fmt.Printf("number of input data %d \n", n)
	fmt.Printf("read line: %s %s %s %s %s-\n", PersonID, LastName, FirstName, Address, City)

	if err != nil {
		log.Fatal(err)
	}
	quety := "Insert into Persons(PersonID, LastName, FirstName, Address, City) Values($1, $2, $3, $4, $5);"

	_, err = db.Exec(quety, PersonID, LastName, FirstName, Address, City)
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
