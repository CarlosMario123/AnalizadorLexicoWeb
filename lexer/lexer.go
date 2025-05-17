package lexer

import (
	"strings"
	"unicode"
)

// Definimos los tipos de tokens
type TokenType int

const (
	IDENTIFICADOR TokenType = iota
	NUMERO
	ERROR
)

// Estructura del token
type Token struct {
	Tipo   TokenType
	Valor  string
	Indice int
}

// Función para analizar léxicamente una cadena
func Lex(input string) []Token {
	var tokens []Token
	words := strings.Fields(input) // Divide la entrada por espacios

	for i, word := range words {
		tokenIndex := i + 1

		if isAllLetters(word) {
			// Es un identificador válido si solo contiene letras
			tokens = append(tokens, Token{Tipo: IDENTIFICADOR, Valor: word, Indice: tokenIndex})
		} else if isAllDigits(word) {
			// Es un número válido si solo contiene dígitos
			tokens = append(tokens, Token{Tipo: NUMERO, Valor: word, Indice: tokenIndex})
		} else {
			// Es un error en cualquier otro caso (identificador inválido)
			tokens = append(tokens, Token{Tipo: ERROR, Valor: word, Indice: tokenIndex})
		}
	}

	return tokens
}

// Función para verificar si una cadena solo contiene letras
func isAllLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// Función para verificar si una cadena solo contiene dígitos
func isAllDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// Función para obtener el nombre del tipo de token como texto
func TokenTypeName(t TokenType) string {
	switch t {
	case IDENTIFICADOR:
		return "IDENTIFICADOR"
	case NUMERO:
		return "NÚMERO"
	case ERROR:
		return "ERROR Identificador inválido (contiene números)"
	default:
		return "DESCONOCIDO"
	}
}