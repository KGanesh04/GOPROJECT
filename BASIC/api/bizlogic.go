package api

import (
	"BASIC/dataservice"
	"database/sql"
	"net/http"
)

func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	//basic validations
	return dataservice.CreateBook(db, w, r)
}
