package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/crud-postgres/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func createConnection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected...")
	return db
}

func GetAllLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Sql Select
	allLanguage, err := GetAllLangage_sql()

	err = json.NewEncoder(w).Encode(allLanguage)
	if err != nil {
		log.Println("Error in encoding in getAlllanguage")
	}
}

func GetByIdLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Unable to convert string to int")
	}

	//Sql Select
	Language, err := GetByIdLanguage_sql(id)

	err = json.NewEncoder(w).Encode(Language)
	if err != nil {
		log.Println("Error in encoding in getAlllanguage")
	}
}

func PostLanguage(w http.ResponseWriter, r *http.Request) {
	var lang models.Language

	err := json.NewDecoder(r.Body).Decode(&lang)
	if err != nil {
		log.Println("Error in decoding in getAllLanguage")
	}

	//Sql Insert
	newLanguage, err := PostLanguage_sql(lang)

	res := response{
		Id:      newLanguage,
		Message: "Add new language",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error in encoding in getAlllanguage")
	}
}

func PutByIdLanguage(w http.ResponseWriter, r *http.Request) {
	var lang models.Language

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println("Unable to convert string to int")
	}

	err = json.NewDecoder(r.Body).Decode(&lang)
	if err != nil {
		log.Println("Error in decoding in getAllLanguage")
	}

	//Sql Insert
	newLanguage, err := PutByIdLanguage_sql(id, lang)

	res := response{
		Id:      newLanguage,
		Message: "Change old language",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error in encoding in getAlllanguage")
	}
}

func DeleteByIdLanguage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Unable to convert string to int")
	}

	//Sql Insert
	newLanguage, err := DeleteByIdLanguage_sql(id)

	res := response{
		Id:      newLanguage,
		Message: "Delte language",
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Error in encoding in getAlllanguage")
	}
}

func GetAllLangage_sql() ([]models.Language, error) {
	db := createConnection()
	defer db.Close()

}

func GetByIdLanguage_sql(id int) (models.Language, error) {
	var lang models.Language
	db := createConnection()
	defer db.Close()
	query := `Select * from LANGUAGES WHERE id=$1`
	row := db.QueryRow(query, id)
	err := row.Scan(&lang.Id, &lang.Year, &lang.Name, &lang.Developer)
	if err != nil {
		log.Println("Unable to execute query GetByIdLanguage_sql")
	}
	return lang, err

}

func PostLanguage_sql(lang models.Language) (int, error) {
	db := createConnection()
	defer db.Close()

	query := `INSERT INTO LANGUAGES (id, year, name, developer) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := db.QueryRow(query, lang).Scan(&id)
	if err != nil {
		log.Println("Unable to execute query PostLanguage_sql")
	}
}

func PutByIdLanguage_sql(id int, lang models.Language) (int, error) {
	db := createConnection()
	defer db.Close()

}

func DeleteByIdLanguage_sql(id int) (int, error) {
	db := createConnection()
	defer db.Close()

}
