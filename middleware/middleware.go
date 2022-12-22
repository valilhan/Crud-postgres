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
	idRowsEffected := PutByIdLanguage_sql(id, lang)

	res := response{
		Id:      idRowsEffected,
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
	idRowsEffected := DeleteByIdLanguage_sql(id)

	res := response{
		Id:      idRowsEffected,
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
	var langs []models.Language
	query := `Select * from LANGUAGES`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Unable to execute query GetAllLangage_sql")
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.Language
		err = rows.Scan(&temp.Id, &temp.Year, &temp.Name, &temp.Developer)
		if err != nil {
			log.Printf("Unable to scan")
		}
		langs = append(langs, temp)
	}
	return langs, nil
}

func GetByIdLanguage_sql(id int) (models.Language, error) {
	var lang models.Language
	db := createConnection()
	defer db.Close()
	query := `Select * from LANGUAGES WHERE id=$1`
	row := db.QueryRow(query, id)
	err := row.Scan(&lang.Id, &lang.Year, &lang.Name, &lang.Developer)
	if err == sql.ErrNoRows {
		log.Printf("No rows with such id")
		return lang, nil
	} else if err != nil {
		return lang, nil
	} else {
		log.Println("Unable to execute query GetByIdLanguage_sql")
		return lang, err
	}
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
	return id, err
}

func PutByIdLanguage_sql(id int, lang models.Language) int {
	db := createConnection()
	defer db.Close()
	query := `UPDATE LANGUAGES SET Year = $1 Name = $2 Developer = $3 WHere Id = %4`
	res, err := db.Exec(query, lang.Year, lang.Name, lang.Developer, lang.Id)
	if err != nil {
		log.Println("Unable to execute query PutByIdLanguage_sql")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Unable to execute query RowsAffected PutByIdLanguage_sql")
	}
	return int(rowsAffected)
}

func DeleteByIdLanguage_sql(id int) int {
	db := createConnection()
	defer db.Close()
	query := `DELETE FROM LANGUAGES WHERE id = %1`
	res, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Unable to execute query DeleteByIdLanguage_sql")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Unable to execute query RowsAffected PutByIdLanguage_sql")
	}
	return int(rowsAffected)
}
