package controllers

import (
	autenticacao "api/src/auth"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("cadastro"); err != nil {
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
	usuarioID, err := repositorio.Criar(usuario)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	usuario.ID = uint(usuarioID)
	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRespositorioDeUsuarios(db)
	usuarios, err := repositorio.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
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
	usuario, err := repositorio.BuscarPorId(usuarioId)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usuario)
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar usuário diferente do seu"))
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = usuario.Preparar("edicao"); err != nil {
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
	if err = repositorio.Atualizar(usuarioID, usuario); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)

}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
	}

	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível Deletar usuário diferente do seu"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
	}
	defer db.Close()
	repositorio := repositorios.NovoRespositorioDeUsuarios(db)
	if err = repositorio.Deletar(usuarioID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	respositorio := repositorios.NovoRespositorioDeUsuarios(db)
	if err = respositorio.Seguir(usuarioID, seguidorID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)

}
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if seguidorID == usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível parar se seguir você mesmo"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	repositorio := repositorios.NovoRespositorioDeUsuarios(db)
	if err = repositorio.DeixarDeSeguir(usuarioID, seguidorID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	respositorio := repositorios.NovoRespositorioDeUsuarios(db)
	seguidores, err := respositorio.BuscarSeguidores(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	respositorio := repositorios.NovoRespositorioDeUsuarios(db)
	seguidores, err := respositorio.BuscarSeguindo(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, seguidores)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de outro usuário"))
		return
	}
	corpoDaRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	var senha modelos.Senha
	if err = json.Unmarshal(corpoDaRequisicao, &senha); err != nil {
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
	senhaSalvaNoBanco, err := repositorio.BuscarSenha(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(senhaSalvaNoBanco, senha.Atual); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, errors.New("a senha não bateru com a do banco"))
		return
	}

	senhaComHash, err := security.Hash(senha.Nova)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)

}
