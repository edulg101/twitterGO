package controllers

import (
	autenticacao "api/src/auth"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {

	corpoDaRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoDaRequisicao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRespositorioDeUsuarios(db)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckPassword(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)

		return
	}
	token, err := autenticacao.CreateToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	usuarioID := strconv.FormatUint(uint64(usuarioSalvoNoBanco.ID), 10)

	respostas.JSON(w,
		http.StatusOK,
		modelos.DadosAutenticacao{
			ID: usuarioID, Token: token,
		},
	)

}
