package mapper

import (
	"yim-go/api/internal/core/domain"
	"yim-go/shared/external/util/typeconvertutil"
	"yim-go/shared/gen/protobuf"
)

func GetBookListToProto(in *domain.GetBookListResponse) *protobuf.GetBookListResponse {
	books := typeconvertutil.ProcessItem(in.Books, bookDomainToProto)
	booksPtr := make([]*protobuf.Book, len(books))
	for i := range books {
		booksPtr[i] = &books[i]
	}

	pagination := &protobuf.PaginationResponse{
		PageId:   int32(in.PaginationResponse.PageId),
		PageSize: int32(in.PaginationResponse.PageSize),
		LastPage: int32(in.PaginationResponse.LastPage),
		TotalElements: int32(in.PaginationResponse.TotalElements),
	}

	return &protobuf.GetBookListResponse{
		Books: booksPtr,
		Pagination: pagination,
	}
}