package main

import (
	"analizadorlexico/lexer" 
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// Estructura para pasar datos a las plantillas
type PageData struct {
	Input   string
	Tokens  []lexer.Token
	Summary struct {
		Identifiers int
		Numbers     int
		Errors      int
	}
}

// Inicializa las plantillas
func init() {
	templatesDir := filepath.Join("templates", "*.html")
	templates = template.Must(template.ParseGlob(templatesDir))
}


func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		
		templates.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		input := r.Form.Get("input")

		
		tokens := lexer.Lex(input)

	
		uniqueIdentifiers := make(map[string]bool)
		uniqueNumbers := make(map[string]bool)
		uniqueErrors := make(map[string]bool)

		for _, token := range tokens {
			switch token.Tipo {
			case lexer.IDENTIFICADOR:
				uniqueIdentifiers[token.Valor] = true
			case lexer.NUMERO:
				uniqueNumbers[token.Valor] = true
			case lexer.ERROR:
				uniqueErrors[token.Valor] = true
			}
		}

		// Preparar datos para la plantilla
		data := PageData{
			Input:  input,
			Tokens: tokens,
		}
		data.Summary.Identifiers = len(uniqueIdentifiers)
		data.Summary.Numbers = len(uniqueNumbers)
		data.Summary.Errors = len(uniqueErrors)

		// Renderizar la plantilla de resultados
		templates.ExecuteTemplate(w, "result.html", data)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}