package main

import (
	"github.com/andreabreu76/estudoGo6/router"
	"github.com/andreabreu76/estudoGo6/utils"
	"log"
	"net/http"
)

func main() {
	r := router.Setup()

	if err := http.ListenAndServe(utils.ServerPort, r); err != nil {
		log.Printf("Erro ao iniciar servico: %v", err)
	}
}
