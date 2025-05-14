package mapper

import (
	"yim-go/api/internal/core/domain"
	"yim-go/shared/external/model/basemodel"
	"yim-go/shared/gen/protobuf"
)

func GetBookListRequestToDomain(in *protobuf.GetBookListRequest) domain.GetBookListRequest {
	var paginationRequest *basemodel.PaginationOffsetRequest
	if in.Pagination != nil {
		paginationRequest = &basemodel.PaginationOffsetRequest{
			PageId: int(in.Pagination.PageId),
			PageSize: int(in.Pagination.PageSize),
		}
	}

	return domain.GetBookListRequest{
		ID: in.Id,
		PaginationRequest: paginationRequest,
	}

}