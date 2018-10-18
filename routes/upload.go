package routes

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/emanpicar/golangUploadImage/app/common"
	"github.com/emanpicar/golangUploadImage/app/config"
	"github.com/emanpicar/golangUploadImage/models"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseMultipartForm(0)
	if common.IsError(w, err, http.StatusInternalServerError) {
		return
	}

	if r.FormValue("token") == config.TOKEN {
		fileName, dataFileBytes, err := getFormFileData(r)
		if common.IsError(w, err, http.StatusInternalServerError) {
			return
		}

		if isValidImage(dataFileBytes) {
			fileSize, err := saveImageToTemp(fileName, dataFileBytes)
			if common.IsError(w, err, http.StatusInternalServerError) {
				return
			}

			err = isImageAllowedToDb(fileName, dataFileBytes, fileSize)
			if common.IsError(w, err, http.StatusInternalServerError) {
				return
			}
		} else {
			if common.IsError(w, errors.New("HTTP Status Forbidden"), http.StatusForbidden) {
				return
			}
		}

	} else {
		if common.IsError(w, errors.New("HTTP Status Forbidden"), http.StatusForbidden) {
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	config.Template.ExecuteTemplate(w, "index.html", config.TOKEN)
}

func getFormFileData(r *http.Request) (string, []byte, error) {
	dataFile, dataHeader, err := r.FormFile("data")
	if err != nil {
		return "", nil, err
	}

	dataFileBytes, err := ioutil.ReadAll(dataFile)
	if err != nil {
		return "", nil, err
	}

	return dataHeader.Filename, dataFileBytes, nil
}

func saveImageToTemp(filename string, dataFileBytes []byte) (int, error) {
	tmpFolder := os.TempDir()

	nf, err := os.Create(filepath.Join(tmpFolder, filename))
	if err != nil {
		return 0, err
	}
	defer nf.Close()

	fileSize, err := nf.Write(dataFileBytes)
	if err != nil {
		return 0, err
	}

	log.Println(fmt.Sprintf("Saving images to %v", tmpFolder))
	return fileSize, nil
}

func isValidImage(dataFileBytes []byte) bool {
	filetype := http.DetectContentType(dataFileBytes)
	isImage := true

	switch filetype {
	case "image/jpeg", "image/jpg":
		log.Println(fmt.Sprintf("%v: valid file type uploaded", filetype))
		break
	case "image/gif":
		log.Println(fmt.Sprintf("%v: valid file type uploaded", filetype))
		break
	case "image/png":
		log.Println(fmt.Sprintf("%v: valid file type uploaded", filetype))
		break
	default:
		log.Println(fmt.Sprintf("%v: Invalid file type uploaded", filetype))
		isImage = false
	}

	return isImage
}

func isImageAllowedToDb(fileName string, blob []byte, fileSize int) error {
	megabyte := 1e+6
	imageSizeLimit := (megabyte * 8)

	if fileSize < int(imageSizeLimit) && fileSize != 0 {
		return models.InsertImageToDB(fileName, blob, fileSize)
	}

	return fmt.Errorf("Image larger than %v bytes", int(imageSizeLimit))
}
