package server

import "net/http"

// NoCache makes browsers revalidate a page on every visit and keeps shared
// caches from storing it.
func NoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Cache-Control", "no-cache, private")
		next.ServeHTTP(resp, req)
	})
}
