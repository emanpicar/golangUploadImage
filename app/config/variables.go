package config

import (
	"log"
	"os"
)

var TOKEN string

func init() {
	TOKEN = os.Getenv("TESTAPP_TOKEN")
	log.Println("TOKEN :##############", TOKEN)
}
