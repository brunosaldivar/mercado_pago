package handlers

import (
	"encoding/json"
	"log"
	helper "mercado_pago/pkg/helper"
	. "mercado_pago/pkg/structs"
	"net/http"

	"github.com/gorilla/mux"
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
func TopSecretSplitGetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	satellites, found := helper.GetCache()

	log.Println("GetCache Top:", satellites)
	if found {
		location, _errCalc := satellites.CalculateCoordinates()
		message, _errMsg := satellites.GetMessage()
		log.Println(_errCalc, _errMsg)
		if _errCalc != nil || _errMsg != nil {
			log.Println(_errCalc, _errMsg)

			json.NewEncoder(w).Encode("No hay suficiente informacion")
		} else {
			var response = &ResponseTopSecret{
				Position: *location,
				Message:  *message,
			}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		json.NewEncoder(w).Encode("Error: No hay suficiente informacion")
	}
}

func TopSecretSplitPostHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//var satellites Satellites
	var satellite Satellite
	vars := mux.Vars(r)
	name := vars["satellite"]

	decodeJson := json.NewDecoder(r.Body)
	err := decodeJson.Decode(&satellite)
	satellite.Name = name

	if err != nil {
		log.Println("Error en payload", err)
		w.WriteHeader(http.StatusNotFound)
	}

	helper.SetCache(satellite)
	w.WriteHeader(http.StatusOK)

}
func ClearCachePostHandler(w http.ResponseWriter, r *http.Request) {
	helper.Clear()
	w.WriteHeader(http.StatusOK)
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
}
