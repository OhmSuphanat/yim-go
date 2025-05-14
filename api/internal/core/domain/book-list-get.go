package domain

import (
	"yim-go/shared/domain"
	"yim-go/shared/external/model/basemodel"
)

type GetBookListRequest struct {
	ID string
	PaginationRequest *basemodel.PaginationOffsetRequest
}

type GetBookListResponse struct {
	Books []domain.Book
	PaginationResponse basemodel.PaginationOffsetResponse
}