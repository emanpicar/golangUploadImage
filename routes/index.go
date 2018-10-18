package routes

import (
	"net/http"

	"github.com/emanpicar/golangUploadImage/app/config"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	config.Template.ExecuteTemplate(w, "index.html", config.TOKEN)
}
