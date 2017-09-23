package handlers

import (
	"net/http"

	"github.com/naoty/refrigerator/handlers/encoders"
)

// RoutesHandler is a HTTP handler to manage HTTP routes.
type RoutesHandler struct {
}

func (h *RoutesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/foods" {
		handler := &FoodsHandler{Encoder: encoders.JSONEncoder{}}
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
