package utils

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

func LoadTemplates(pattern string) {
	Templates = template.Must(template.ParseGlob(pattern))
}

func Load2Templates(template1 template.Template, template2 template.Template){
	template.ParseFiles()
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	Templates.ExecuteTemplate(w, tmpl, data)
}
