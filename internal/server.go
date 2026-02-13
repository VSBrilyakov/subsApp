package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/VSBrilyakov/subsApp/configs"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config configs.ServerConfig, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.GetAddress(),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
