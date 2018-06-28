package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"simpleBidParser/routes"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", routes.IndexHandler).Methods("POST","GET")
	router.NotFoundHandler = http.HandlerFunc(routes.NotFoundHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
