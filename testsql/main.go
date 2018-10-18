package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

type Images struct {
	Id       string
	FileName string
	Blob     string
	FileSize float32
}

func main() {
	var dbhost = os.Getenv("TESTAPP_DBHOST")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:password@%v/postgres?sslmode=disable", dbhost))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("You are connected to the database")

	rows, err := db.Query("SELECT * FROM images;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	imgs := make([]Images, 0)
	for rows.Next() {
		img := Images{}
		err := rows.Scan(&img.Id, &img.FileName, &img.Blob, &img.FileSize)
		if err != nil {
			panic(err)
		}
		imgs = append(imgs, img)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, img := range imgs {
		// fmt.Println(bk.isbn, bk.title, bk.author, bk.price)
		fmt.Printf("%s, %s, $%.2f\n", img.Id, img.FileName, img.FileSize)

		nf, err := os.Create(filepath.Join("./img/", img.FileName))
		if err != nil {
			panic(err)
		}
		defer nf.Close()

		fileSize, err := nf.Write([]byte(img.Blob))
		if err != nil {
			panic(err)
		}

		log.Println(fileSize, "######", int(img.FileSize))
	}
}
