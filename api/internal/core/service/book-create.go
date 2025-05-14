package service

import (
	"context"
	"time"
	"yim-go/api/internal/core/domain"
)

func (s *service) CreateBook(ctx context.Context, in domain.CreateBookRequest) error {
	in.CreatedAt  = time.Now()
	in.UpdatedAt  = in.CreatedAt

	if err := s.repo.CreateBook(ctx, in); err != nil {
		return err
	}

	return nil

}