package http

import (
	// stdlib
	"net/http"
	"log"
	// universe
	"github.com/universelabs/universe-core/universe"
)

// HTTP Service
type Server struct {
	// net/http infrastructure
	ln net.Listener
	// handler to serve
	Handler *Handler
	// 
	Addr string
}

// Returns a new Server instance
func NewServer(addr string, km universe.KeyManager) *Server {
	srv := &Server{
		// init handler
		Handler: NewHandler(km),
		// set addr
		Addr: addr,	
	}
	return srv
}

// Listens to addr and serves from s.Handler. http.Serve is called in a
// goroutine, so main() must hang for the process not to close!
func (s *Server) Open() error {
	// open the socket
	ln, err := ln.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	// serve http
	go func() { log.Fatal(http.Serve(s.ln, s.Handler)) }()

	return nil
}

// Close the server
func (s *Server) Close() error {
	if s.ln != nil {
		return s.ln.Close()
	}
	return nil
}
