package handlers

import (
	"log"
	"net/http"
	"os"
)

// LoggingHandler is a HTTP handler to log accesses.
type LoggingHandler struct {
	handler http.Handler

	// Logger is a logger to log accesses.
	logger *log.Logger
}

// Log returns a new LoggingHandler wrapping given HTTP handler.
func Log(h http.Handler) *LoggingHandler {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	return &LoggingHandler{
		handler: h,
		logger:  logger,
	}
}

func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
	h.logger.Printf("%s %s\n", r.Method, r.URL)
}
