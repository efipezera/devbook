package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//CarregarTemplates insere os templates HTML na variável templates.
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

//ExecutarTemplate renderiza uma página HTML na tela.
func ExecutarTemplate(rw http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(rw, template, dados)
}
