package config

import (
	"html/template"
)

var Template *template.Template

func init() {
	Template = template.Must(template.ParseFiles("index.html"))
}
