package main

import (
	"context"
	"log/slog"
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/contract"
)

type Server struct {
	logger   *slog.Logger
	db       contract.Storage
	settings *settings.ServerSettings
}

func NewServer(s *settings.ServerSettings, logger *slog.Logger, storage contract.Storage) (*Server, error) {
	// в зависимости от конфига создаем сервер нужного типа со всеми его зависимостями
	return &Server{logger: logger, settings: s, db: storage}, nil
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
