package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Render's templates using html/templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./internal/templates/" + tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
