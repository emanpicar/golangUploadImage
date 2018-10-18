package common

import (
	"fmt"
	"log"
	"net/http"
)

func IsError(w http.ResponseWriter, err error, httpError int) bool {
	if err != nil {
		log.Println("Server error...", fmt.Sprintf("Error %v: %v", httpError, err))
		http.Error(w, fmt.Sprintf("Error %v: %v", httpError, err), httpError)
		return true
	}

	return false
}
