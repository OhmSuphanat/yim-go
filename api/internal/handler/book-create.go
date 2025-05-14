package handler

import (
	"context"
	"yim-go/api/internal/mapper"
	"yim-go/shared/gen/protobuf"

	"google.golang.org/protobuf/types/known/emptypb"
)



func (h *handler) CreateBook(ctx context.Context, req *protobuf.CreateBookRequest) (*emptypb.Empty, error) {
	if err := h.svc.CreateBook(ctx, mapper.CreateBookRequestToDomain(req)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}