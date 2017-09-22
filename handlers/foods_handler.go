package handlers

import (
	"fmt"
	"net/http"
)

// FoodsHandler is a HTTP handler for foods.
type FoodsHandler struct {
}

func (h *FoodsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: Return foods")
	w.WriteHeader(http.StatusOK)
}
