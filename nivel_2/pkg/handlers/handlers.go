package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/brunosaldivar/mercado_pago/nivel_2/pkg/structs"
)

func TopSecretPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var satellites Satellites
	decodeJson := json.NewDecoder(r.Body)
	err := decodeJson.Decode(&satellites)

	location, _errCalc := satellites.CalculateCoordinates()
	message, _errMsg := satellites.GetMessage()

	if err != nil {
		log.Println("Error en payload", err)
		w.WriteHeader(http.StatusNotFound)
	}
	if _errCalc != nil || _errMsg != nil {
		log.Println(_errCalc, _errMsg)
		w.WriteHeader(http.StatusNotFound)
	} else {
		var response = &ResponseTopSecret{
			Position: *location,
			Message:  *message,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
}
