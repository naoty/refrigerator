package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naoty/refrigerator/handlers"
	"github.com/naoty/refrigerator/handlers/encoders"
	"github.com/naoty/refrigerator/infra"
)

func main() {
	config := infra.DefaultConfig
	log.Printf("user: %v, pass: %v", config.DBUser, config.DBPassword)

	routesHandler := &handlers.RoutesHandler{}
	routesHandler.GET("/foods", &handlers.FoodsHandler{Encoder: encoders.JSONEncoder{}})

	handler := handlers.Log(routesHandler)
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
