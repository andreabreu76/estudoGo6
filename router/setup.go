package router

import (
	"github.com/andreabreu76/estudoGo6/handlers"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	router := mux.NewRouter()

	prefixVersion := router.PathPrefix("/api/v1").Subrouter()

	prefixVersion.HandleFunc("/", handlers.Index).Methods("GET")

	prefixVersion.HandleFunc("/cep/{cep}", handlers.ConsultaCep).Methods("GET")

	return router
}
