package server

import (
	"context"
	"io"
	"net/http"

	"example.org/log"
)

type Server struct {
	logger *log.Logger
	addr   string
}

func NewServer(logger *log.Logger, addr string) *Server {
	return &Server{
		logger: logger,
		addr:   addr,
	}
}

func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.logger.Debug("Receiver request")
		io.WriteString(w, "Hello")
	})
	return nil
}
