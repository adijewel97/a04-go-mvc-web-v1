package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "USERADISMONLAP/adis123@127.0.0.1:1521/adis.iconpln.co.id"
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable: ", err)
	}

	fmt.Println("Connected to Oracle Database!")
	DB = db
}
