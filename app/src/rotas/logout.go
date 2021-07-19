package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogout = []Rota{
	{
		URI:                "/logout",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FazerLogout,
		RequerAutenticacao: true,
	},
}
