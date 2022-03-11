package utils

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

func LoadTemplates(pattern string) {
	Templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	Templates.ExecuteTemplate(w, tmpl, data)
}
