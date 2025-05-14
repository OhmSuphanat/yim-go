package port

import (
	"context"
	"yim-go/api/internal/core/domain"
)

type Service interface {
	CreateBook(ctx context.Context, in domain.CreateBookRequest) error
	GetBookList(ctx context.Context, in domain.GetBookListRequest) (*domain.GetBookListResponse, error)
}
