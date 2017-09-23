package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naoty/refrigerator/handlers"
	"github.com/naoty/refrigerator/handlers/encoders"
)

func main() {
	http.Handle("/foods", handlers.Log(&handlers.FoodsHandler{Encoder: encoders.JSONEncoder{}}))
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
