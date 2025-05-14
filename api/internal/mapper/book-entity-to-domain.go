package mapper

import (
	"yim-go/shared/domain"
	"yim-go/shared/entity"
)

func BookEntityToDomain(e entity.Book) domain.Book {
	return domain.Book{
		Id: e.Id,
		Title: e.Title,
		Author: e.Author,
		Description: e.Description,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}