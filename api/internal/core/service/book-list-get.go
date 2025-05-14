package service

import (
	"context"
	"fmt"
	"yim-go/api/internal/core/domain"
)

func (s *service) GetBookList(ctx context.Context, in domain.GetBookListRequest) (*domain.GetBookListResponse, error) {
	getBookListResp, err := s.repo.GetBookList(ctx, in)
	if err != nil {
		fmt.Printf("unable to get book list: %+v\n", err)
		return nil, err
	}
	return getBookListResp, nil
}