package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

//FazerLogout remove os dados de autenticação salvos no browser do usuário.
func FazerLogout(rw http.ResponseWriter, r *http.Request) {
	cookies.Deletar(rw)
	http.Redirect(rw, r, "/login", 302)
}
