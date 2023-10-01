package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server with HTTP protocol
type Server struct {
	server *http.Server
}

// New http server
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

// Start http server
func (s *Server) Start() (err error) {
	logrus.Printf("Server started at:  %s", s.server.Addr)
	go func() {
		err = s.server.ListenAndServe()
	}()

	return
}

// Shutdown stoped http server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
