package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	chave := make([]byte, 64)
// 	if _, err := rand.Read(chave); err != nil {
// 		panic(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {

	config.Carregar()

	fmt.Printf("Escutando na porta %d \n", config.Port)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
