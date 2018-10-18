package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	var dbhost = os.Getenv("TESTAPP_DBHOST")
	DB, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:password@%v/postgres?sslmode=disable", dbhost))
	if err != nil {
		log.Fatalln(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln(err)
	}

	createImagesTable()

	log.Println("Database connection successful...")
}

func createImagesTable() {
	_, err := DB.Exec(
		`CREATE TABLE IF NOT EXISTS images (
			id SERIAL PRIMARY KEY, 
			file_name text, 
			blob bytea, 
			file_size bigint
		);`,
	)

	if err != nil {
		log.Fatalln(err)
	}
}
