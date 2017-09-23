package handlers

import (
	"log"
	"net/http"
	"os"
	"time"
)

// LoggingHandler is a HTTP handler to log accesses.
type LoggingHandler struct {
	handler http.Handler

	// Logger is a logger to log accesses.
	logger *log.Logger
}

// Log returns a new LoggingHandler wrapping given HTTP handler.
func Log(h http.Handler) *LoggingHandler {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	return &LoggingHandler{
		handler: h,
		logger:  logger,
	}
}

func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now().UnixNano()
	h.handler.ServeHTTP(w, r)
	finishTime := time.Now().UnixNano()
	responseTimeMicro := (finishTime - startTime) / 1000
	h.logger.Printf("%s %s %d µs\n", r.Method, r.URL, responseTimeMicro)
}
