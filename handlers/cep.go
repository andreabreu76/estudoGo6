package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreabreu76/estudoGo6/entities"
	"github.com/andreabreu76/estudoGo6/utils"
	"github.com/gorilla/mux"
)

func ConsultaCep(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)

	if request["cep"] == "" {
		log.Println("CEP nao informado")
		utils.JsonResponse(w, http.StatusBadRequest, "CEP nao informado")
		return
	}

	apiURL := "https://ws.apicep.com/cep.json?code=" + request["cep"]

	res, err := http.Get(apiURL)
	if err != nil {
		log.Println("Erro ao consultar CEP: ", err)
		utils.JsonResponse(w, http.StatusInternalServerError, "Erro ao consultar CEP")
		return
	}

	defer res.Body.Close()

	var cepresponse entities.CepResponse

	if err := json.NewDecoder(res.Body).Decode(&cepresponse); err != nil {
		log.Println("Erro ao decodificar JSON: ", err)
		utils.JsonResponse(w, http.StatusInternalServerError, "Erro ao decodificar JSON")
		return
	}

	if !cepresponse.OK {
		log.Println("CEP nao encontrado: ", cepresponse.StatusText)
		utils.JsonResponse(w, http.StatusBadRequest, "Erro ao consultar CEP")
		return
	}

	utils.JsonResponse(w, http.StatusOK, cepresponse)

}
