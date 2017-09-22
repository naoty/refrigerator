package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naoty/refrigerator/models"
)

// FoodsHandler is a HTTP handler for foods.
type FoodsHandler struct {
}

func (h *FoodsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get data from database
	foods := []models.Food{
		models.Food{Name: "Apple", Quantity: models.Quantity{Value: 1}},
		models.Food{Name: "Orange", Quantity: models.Quantity{Value: 2}},
	}
	data, err := json.Marshal(foods)
	if err != nil {
		fmt.Println("TODO: error handling")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
