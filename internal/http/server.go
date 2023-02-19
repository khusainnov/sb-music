package http

import (
	"net/http"

	"github.com/khusainnov/sb-music/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg *config.Config
	srv *http.Server
}

func New(cfg *config.Config) *Server {
	srv := &http.Server{
		Addr: cfg.HTTPAddr,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}
}

func (s *Server) RunHTTP() {
	s.srv.Handler = s.initEndpoints()

	go func() {
		s.cfg.Logger.Info("start listening http", zap.String("PORT", s.cfg.HTTPAddr))
		if err := s.srv.ListenAndServe(); err != nil {
			s.cfg.Logger.Fatal("error run http", zap.Error(err))
		}
	}()
}
