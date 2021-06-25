package middlewares

import (
	autenticacao "api/src/auth"
	"api/src/respostas"
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s, %s, %s", r.RequestURI, r.Method, r.Host)
		next(w, r)
	}

}

// Autenticar verifica se o usuario fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("Autenticando ..... ")
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
