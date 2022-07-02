package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"cs101/pkg/telementry/logging"
)

type Server struct {
	cfg Config

	logger logging.Logger
	http   *http.Server
	db     *sql.DB
}

func New(c Config) *Server {
	return &Server{
		cfg: c,
	}
}

func (s *Server) Run() {
	s.init()
	s.start()
}

func (s *Server) start() {
	go func() {
		s.logger.Infof("server start HTTP on %s", s.http.Addr)
		if err := s.startHTTP(); err != nil {
			s.logger.Fatalf("failed to start HTTP server: %+v", err)
		}
	}()
}

func (s *Server) startHTTP() error {
	if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s Server) Shutdown() {
	fmt.Println("Server is shutting down...")

	fmt.Println("Server shutdown")
}
