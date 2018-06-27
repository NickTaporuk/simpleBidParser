package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"simpleBidParser/routes"
)

func main() {
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.NotFound = http.HandlerFunc(routes.NotFound)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
