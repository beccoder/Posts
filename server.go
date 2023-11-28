package Blogs

import (
	"context"
	"github.com/rs/cors"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(httpAddr string, handler http.Handler) error {
	handle := cors.AllowAll().Handler(handler)
	s.httpServer = &http.Server{
		Addr:           httpAddr,
		Handler:        handle,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
