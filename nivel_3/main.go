package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "mercado_pago/nivel_3/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/topsecret", handlers.TopSecretPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite}", handlers.TopSecretSplitPostHandler).Methods("POST")
	router.HandleFunc("/topsecret_split/", handlers.TopSecretSplitGetHandler).Methods("GET")
	router.HandleFunc("/clear", handlers.ClearCachePostHandler).Methods("POST")
	log.Printf("Started on port:", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("Failed starting http server: ", err)
	}

}
