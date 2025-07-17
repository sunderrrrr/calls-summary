package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   5 * time.Second,
		ReadTimeout:    5 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
