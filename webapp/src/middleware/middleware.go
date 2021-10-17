package middleware

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

//Logger escreve informações da requisição no console.
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(rw, r)
	}
}

//Autenticar verifica a existência de cookies.
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(rw, r, "/login", 302)
			return
		}
		proximaFuncao(rw, r)
	}
}
