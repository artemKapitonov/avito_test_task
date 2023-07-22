package httpserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server *http.Server
}

// New -.
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

func (s *Server) Start() error {
	logrus.Printf("Server start on Addr:  %s", s.server.Addr)
	return s.server.ListenAndServe()
}
