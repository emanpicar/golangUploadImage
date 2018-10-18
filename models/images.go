package models

import (
	"fmt"
	"log"
)

func InsertImageToDB(fileName string, blob []byte, fileSize int) error {
	_, err := DB.Exec(
		`INSERT INTO images (
			file_name, 
			blob, 
			file_size
		) VALUES ($1, $2, $3)`,
		fileName,
		blob,
		fileSize,
	)

	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Inserting %v to DB successful", fileName))
	return nil
}
