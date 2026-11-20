package server

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Health returns a handler reporting that the server is up.
func Health() http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		logrus.WithFields(logrus.Fields{"method": req.Method, "path": req.URL.Path, "remote": req.RemoteAddr}).Info("health check")
		resp.Header().Set("Content-Type", "application/json")
		_, _ = resp.Write([]byte(`{"status":"ok"}`))
	})
}

// Index returns a handler rendering the home page. The page is rendered into a
// buffer first, so a template error gives a clean 500 and not a half-sent page.
func Index(pages *template.Template) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, _ *http.Request) {
		var page bytes.Buffer
		if err := pages.ExecuteTemplate(&page, "index.html", nil); err != nil {
			logrus.WithError(err).Error("rendering the home page")
			http.Error(resp, "internal server error", http.StatusInternalServerError)

			return
		}
		resp.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = page.WriteTo(resp)
	})
}
