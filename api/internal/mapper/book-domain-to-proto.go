package mapper

import (
	"strconv"
	"yim-go/shared/domain"
	"yim-go/shared/gen/protobuf"
)

func bookDomainToProto(in domain.Book) protobuf.Book {
	return protobuf.Book{
		Id: strconv.FormatInt(in.Id, 10),
		Title: in.Title,
		Author: in.Author,
		Description: in.Description,
		CreatedAt: in.CreatedAt.String(),
		UpdatedAt: in.UpdatedAt.String(),
	}
}