package handlers

import (
	"net/http"
)

// RoutesHandler is a HTTP handler to manage HTTP routes.
type RoutesHandler struct {
	routes map[string]map[string]http.Handler
}

// GET adds a new route to a given path in GET method.
func (h *RoutesHandler) GET(path string, handler http.Handler) {
	h.register("GET", path, handler)
}

// POST adds a new route to a given path to POST method.
func (h *RoutesHandler) POST(path string, handler http.Handler) {
	h.register("POST", path, handler)
}

func (h *RoutesHandler) register(method, path string, handler http.Handler) {
	if h.routes == nil {
		h.routes = make(map[string]map[string]http.Handler)
	}

	_, ok := h.routes[method]
	if !ok {
		h.routes[method] = make(map[string]http.Handler)
	}
	h.routes[method][path] = handler
}

func (h *RoutesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths, ok := h.routes[r.Method]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler, ok := paths[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler.ServeHTTP(w, r)
}
