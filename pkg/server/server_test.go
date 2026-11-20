package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/lamboley/wowhabits/pkg/webstatic"
)

func newTestServer(t *testing.T) *Server {
	t.Helper()

	srv, err := New(Config{Addr: "127.0.0.1:0"})
	if err != nil {
		t.Fatal(err)
	}

	return srv
}

func serve(handler http.Handler, method, target string, headers ...string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, nil)
	for i := 0; i+1 < len(headers); i += 2 {
		req.Header.Set(headers[i], headers[i+1])
	}
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	return rec
}

func TestHealth(t *testing.T) {
	handler := newTestServer(t).Handler()

	rec := serve(handler, http.MethodGet, "/health")
	if rec.Code != http.StatusOK || rec.Body.String() != `{"status":"ok"}` {
		t.Fatalf("status %d, body %q", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("Content-Type = %q", got)
	}
	if got := serve(handler, http.MethodPost, "/health").Code; got != http.StatusMethodNotAllowed {
		t.Fatalf("POST /health: status %d, want 405", got)
	}
}

func TestIndex(t *testing.T) {
	rec := serve(newTestServer(t).Handler(), http.MethodGet, "/")

	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("Content-Type = %q", got)
	}
	if got := rec.Header().Get("Cache-Control"); got != "no-cache, private" {
		t.Fatalf("Cache-Control = %q, want no-cache, private", got)
	}
	if !strings.Contains(rec.Body.String(), webstatic.URL("htmx.min.js")) {
		t.Fatalf("the page does not reference the hashed htmx url:\n%s", rec.Body.String())
	}
}

func TestOnlyTheRootIsTheIndex(t *testing.T) {
	if got := serve(newTestServer(t).Handler(), http.MethodGet, "/nope").Code; got != http.StatusNotFound {
		t.Fatalf("GET /nope: status %d, want 404", got)
	}
}

func TestStaticAsset(t *testing.T) {
	handler := newTestServer(t).Handler()

	rec := serve(handler, http.MethodGet, webstatic.URL("htmx.min.js"))
	if rec.Code != http.StatusOK || rec.Body.Len() == 0 {
		t.Fatalf("status %d, %d bytes", rec.Code, rec.Body.Len())
	}
	if got := rec.Header().Get("Cache-Control"); !strings.Contains(got, "max-age=31536000") {
		t.Fatalf("Cache-Control = %q, want a one year max-age", got)
	}

	etag := rec.Header().Get("ETag")
	if got := serve(handler, http.MethodGet, webstatic.URL("htmx.min.js"), "If-None-Match", etag).Code; got != http.StatusNotModified {
		t.Fatalf("conditional request: status %d, want 304", got)
	}
}

func TestStartStop(t *testing.T) {
	srv := newTestServer(t)
	done := make(chan error, 1)
	go func() { done <- srv.Start() }()

	if err := srv.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Start returned %v after a graceful stop", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Start did not return after Stop")
	}
}

func TestStartFailsOnBadAddress(t *testing.T) {
	srv, err := New(Config{Addr: "bad:addr:x"})
	if err != nil {
		t.Fatal(err)
	}
	if err := srv.Start(); err == nil {
		t.Fatal("Start should fail on an invalid address")
	}
}
