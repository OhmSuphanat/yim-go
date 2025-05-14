package service

import "yim-go/api/internal/core/port"

type service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &service{
		repo: repo,
	}
}
