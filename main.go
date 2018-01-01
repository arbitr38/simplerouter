package main

import (
	"log"
	"net/http"

	"github.com/arbitr38/simplerouter/handlers"
)

func main() {
	log.Print("Startin the service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))

}
