package main

import (
	"encoding/json"
	"log"
	"net/http"

	structs "github.com/brunosaldivar/mercado_pago/nivel_2/pkg"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/topsecret", topSecretHandler).Methods("POST")
	router.HandleFunc("/topsecret_split", topSecretSplitPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split", topSecretHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))
}

func topSecretSplitPostHandler(w http.ResponseWriter, r *http.Request) {

	var satellites structs.Satellites
	decodeJson := json.NewDecoder(r.Body)
	err := decodeJson.Decode(&satellites)
	//TODO: len(satellites) < 3 datos incompletos

	location, _err := satellites.CalculateCoordinates()
	message := satellites.GetMessage()

	if err != nil {
		panic(err) // TODO: no salir con este error
	}
	if _err != nil {
		panic(_err) // TODO: no salir con este error
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structs.ResponseTopSecret{
		Position: location,
		Message:  message,
	})
	//TODO: 200 Y ERROR 404
}
func topSecretHandler(w http.ResponseWriter, r *http.Request) {

	var satellites structs.Satellites
	decodeJson := json.NewDecoder(r.Body)
	err := decodeJson.Decode(&satellites)

	//TODO: len(satellites) < 3 datos incompletos

	location, _err := satellites.CalculateCoordinates()
	message := satellites.GetMessage()

	if err != nil {
		panic(err) // TODO: no salir con este error
	}
	if _err != nil {
		panic(_err) // TODO: no salir con este error
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structs.ResponseTopSecret{
		Position: location,
		Message:  message,
	})
	//TODO: 200 Y ERROR 404
	//respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}
