package handlers

import (
	"fmt"
	"net/http"

	"github.com/naoty/refrigerator/handlers/encoders"
	"github.com/naoty/refrigerator/repositories"
)

// FoodsHandler is a HTTP handler for foods.
type FoodsHandler struct {
	// Encoder is used to encode HTTP responses.
	Encoder encoders.Encoder
}

func (h *FoodsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	repo := repositories.FoodsRepository{}
	foods := repo.GetFoods()
	data, err := h.Encoder.Encode(foods)
	if err != nil {
		fmt.Println("TODO: error handling")
	}

	w.Header().Set("Content-Type", h.Encoder.ContentType())
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
