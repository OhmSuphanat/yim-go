package handler

import (
	"context"
	"yim-go/api/internal/mapper"
	"yim-go/shared/gen/protobuf"
)

func (h *handler) GetBookList(ctx context.Context, req *protobuf.GetBookListRequest) (*protobuf.GetBookListResponse, error) {
	resp, err := h.svc.GetBookList(ctx, mapper.GetBookListRequestToDomain(req))
	if err != nil {
		return nil, err
	}
	return mapper.GetBookListToProto(resp), nil
	
}