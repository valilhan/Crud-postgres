package router

import (
	"github/crud-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/languages", middleware.getAllLanguage()).Methods("GET")             // Get all languages
	router.HandleFunc("/languages/{id}", middleware.GetByIdLanguage()).Methods("GET")       // Get by Id a language
	router.HandleFunc("/languages", middleware.PostLanguage()).Methods("POST")              // Create new language by Id a language
	router.HandleFunc("/languages/{id}", middleware.PutByIdLanguage()).Methods("PUT")       // Change language by Id
	router.HandleFunc("/languages/{id}", middleware.DeleteByIdLanguage()).Methods("DELETE") // Delete language by Id

}
