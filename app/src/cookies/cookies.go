package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	dadosCodificados, err := s.Encode("dados", dados)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

func Ler(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("dados")
	if err != nil {
		return nil, err
	}
	valores := make(map[string]string)
	if err = s.Decode("dados", cookie.Value, &valores); err != nil {
		return nil, err
	}
	return valores, nil
}

func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
