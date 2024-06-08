package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
)

// Server with HTTP protocol.
type Server struct {
	log    *slog.Logger
	server *http.Server
}

// New is creating new http server.
func New(handler http.Handler, port string, log *slog.Logger) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         ":" + port,
	}

	s := &Server{
		log:    log,
		server: httpServer,
	}

	return s
}

// Start is starting http server.
func (s *Server) Start() {
	const op = "server.Start"

	log := s.log.With(slog.String("op", op))

	g := new(errgroup.Group)

	log.Info("Server started at:" + s.server.Addr)

	g.Go(s.server.ListenAndServe)
}

// Shutdown id stopping http server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
