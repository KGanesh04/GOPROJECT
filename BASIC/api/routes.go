package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/create", CreateHandle(db))

}

// post=>create a record in DB
//put=>update a record in DB
//delete=>delete a record in DB
//get=>get a record in DB
