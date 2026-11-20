package server

import "net/http"

// Config is the root config for Server.
type Config struct {
	Addr string
}

// Server is the root datastructure for Server.
type Server struct {
	cfg Config
	srv *http.Server
}
