package templates

import (
	"net/http"
	"text/template"
)

// Renders the base.html template plus a content template.
func Render(w http.ResponseWriter, contentTemplate string, data interface{}) error {
	templates := template.Must(template.ParseFiles("templates/base.html", "templates/"+contentTemplate))
	return templates.ExecuteTemplate(w, "base.html", data)
}
