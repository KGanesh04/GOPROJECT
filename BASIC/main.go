package main

import (
	"BASIC/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:G@ni053012@tcp(127.0.0.1:3306)/gocrud?parseTime=true" // Intilizing databse connection

	db, err := sql.Open("mysql", dsn) // preparing the database drivers and validate the arugumets passing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // closing the database connection

	if err := db.Ping(); err != nil { // checking the database connection
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully") // printing the database connected message(Terminal view)
	api.RegisterRoutes(db)                         // Registering the routes
	fmt.Println("Sent the request to routers")     // printing the server started message(Terminal view)

	log.Println("server strated on port 8080")   // printing the server started message(Terminal view)
	log.Fatal(http.ListenAndServe(":8080", nil)) // starting the server on port 8080

}
