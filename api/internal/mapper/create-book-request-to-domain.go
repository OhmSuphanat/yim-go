package mapper

import (
	"yim-go/api/internal/core/domain"
	"yim-go/shared/gen/protobuf"
)

func CreateBookRequestToDomain(req *protobuf.CreateBookRequest) domain.CreateBookRequest {
	return domain.CreateBookRequest{
		Title:       req.Title,
		Author: 	 req.Author,
		Description: req.Description,
	}
}