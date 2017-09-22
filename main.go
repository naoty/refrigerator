package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/naoty/refrigerator/handlers"
	"github.com/naoty/refrigerator/handlers/encoders"
)

func main() {
	fmt.Println("Hello, world!")

	http.Handle("/foods", &handlers.FoodsHandler{Encoder: encoders.JSONEncoder{}})
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
