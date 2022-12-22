package main

import (
	"github/crud-postgres/router"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	router := router.Router()
	log.Println("Server starting...")

	log.Fatal(http.ListenAndServe(":8000", router))

}
