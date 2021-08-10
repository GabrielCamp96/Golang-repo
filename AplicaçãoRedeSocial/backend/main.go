package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	log.Println(fmt.Sprintf("Rodando API na porta %d!", config.ApiPort))

	r := router.Create()

	http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r)
}
