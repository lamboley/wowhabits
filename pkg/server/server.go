package server

import (
	"context"
	"embed"
	"errors"
	"html/template"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/lamboley/wowhabits/pkg/webstatic"
)

//go:embed templates/*.html
var templatesFS embed.FS

// New returns a Server built from cfg. It does not listen until Start is
// called. The templates are parsed here, so a broken template fails at startup
// and not on the first request.
func New(cfg Config) (*Server, error) {
	pages, err := template.New("").Funcs(webstatic.Funcs()).ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("GET "+webstatic.URLPrefix+"/", webstatic.Handler())
	mux.Handle("GET /health", Health())
	mux.Handle("GET /{$}", NoCache(Index(pages)))

	return &Server{
		cfg: cfg,
		srv: &http.Server{
			Addr:              cfg.Addr,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}, nil
}

// Handler returns the HTTP handler of the server, to serve requests without
// listening on a port.
func (s *Server) Handler() http.Handler {
	return s.srv.Handler
}

// Start listens on the configured address and serves requests until Stop is
// called. It returns nil after a graceful stop.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		return err
	}

	logrus.Infof("listening on %s", listener.Addr())
	if err := s.srv.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Stop shuts the server down gracefully, waiting for in-flight requests until
// ctx is done.
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
