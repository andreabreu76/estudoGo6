package handlers

import (
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"message\":\"Olá Mundo!\"}"))
}
