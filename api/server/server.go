package server

import (
	"context"
	"golvl2/app/domain"
	"net/http"
	"time"
)

type Server struct {
	srv   http.Server
	repos domain.Repositories
}

func (s *Server) Start(r domain.Repositories) {
	s.repos = r
	go s.srv.ListenAndServe()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}
	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	return s
}
