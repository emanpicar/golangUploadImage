package main

import (
	"net/http"

	"github.com/emanpicar/golangUploadImage/routes"
)

func main() {
	http.HandleFunc("/", routes.IndexHandler)
	http.HandleFunc("/upload", routes.UploadHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
