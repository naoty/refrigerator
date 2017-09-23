package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naoty/refrigerator/handlers"
)

func main() {
	handler := handlers.Log(&handlers.RoutesHandler{})
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
