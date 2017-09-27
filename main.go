package main

import (
	"log"
	"net/http"
	"os"

	"github.com/naoty/refrigerator/config"
	"github.com/naoty/refrigerator/handlers"
	"github.com/naoty/refrigerator/handlers/encoders"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user: %v, pass: %v", c.DBUser, c.DBPassword)

	routesHandler := &handlers.RoutesHandler{}
	routesHandler.GET("/foods", &handlers.FoodsHandler{Encoder: encoders.JSONEncoder{}})

	handler := handlers.Log(routesHandler)
	err = http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
