package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ✅ Global Database Instance
var DB *sqlx.DB

// ✅ Connect to Database
func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("mysql", "root:G@ni053012@tcp(127.0.0.1:3306)/gocrud")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	fmt.Println("✅ Database connected successfully!")
}
