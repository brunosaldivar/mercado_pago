package main

import (
	"log"
	"net/http"

	handlers "github.com/brunosaldivar/mercado_pago/nivel_2/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/topsecret", handlers.TopSecretPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split", handlers.TopSecretSplitPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split", handlers.TopSecretSplitGetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))

}
