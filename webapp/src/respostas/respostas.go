package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//ErroAPI representa a resposta de erro da API.
type ErroAPI struct {
	Erro string `json:"erro"`
}

//JSON retorna uma resposta em formato JSON para a requisição.
func JSON(rw http.ResponseWriter, statusCode int, dados interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(rw).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

//TratarStatusCodeDeErro trata as requisições com status code 400 ou superior.
func TratarStatusCodeDeErro(rw http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(rw, r.StatusCode, erro)
}
