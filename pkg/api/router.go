// Package api is the HTTP REST adapter.
package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// NewRouter returns the HTTP handler serving the REST API.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"method": r.Method, "path": r.URL.Path, "remote": r.RemoteAddr}).Info("health check")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	return mux
}
