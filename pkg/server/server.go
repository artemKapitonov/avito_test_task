package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
)

// Server with HTTP protocol.
type Server struct {
	server *http.Server
}

// New is creating new http server.
func New(handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         ":" + port,
	}

	s := &Server{
		server: httpServer,
	}

	return s
}

// Start is starting http server.
func (s *Server) Start() error {
	var err error

	slog.Info(fmt.Sprintf("Server started at: %s", s.server.Addr))

	go func() {
		err = s.server.ListenAndServe()
	}()

	return err
}

// Shutdown id stopping http server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
