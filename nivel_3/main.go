package main

import (
	"log"
	"net/http"

	handlers "github.com/brunosaldivar/mercado_pago/nivel_3/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/topsecret", handlers.TopSecretPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite}", handlers.TopSecretSplitPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split/", handlers.TopSecretSplitGetHandler).Methods("GET")
	router.HandleFunc("/clear", handlers.ClearCachePostHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", router))

}
