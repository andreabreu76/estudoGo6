package handlers

import (
	"net/http"

	"github.com/andreabreu76/estudoGo6/utils"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	utils.JsonResponse(w, http.StatusOK, "API está funcionando")
}
