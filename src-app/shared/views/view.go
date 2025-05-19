package view

import (
	"embed"
	"net/http"
	"text/template"
)

// Intilizaing HTML Templates to parse based on requests
var templates *template.Template

func Parse(t embed.FS) {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// Render function to render HTML templates with data
func Render(w http.ResponseWriter, filename string, data any) {
	err := templates.ExecuteTemplate(w, filename, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
