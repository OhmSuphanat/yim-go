package repository

import (
	"context"
	"yim-go/api/internal/core/domain"
	"yim-go/api/internal/mapper"
	"yim-go/shared/constant"
	"yim-go/shared/entity"
	"yim-go/shared/external/model/basemodel"
	"yim-go/shared/external/util/typeconvertutil"

	"github.com/huandu/go-sqlbuilder"
)

func setWhere(in domain.GetBookListRequest, sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	filters := map[string]any{
		"id": func() any{
			if in.ID == "" {
				return nil
			}
			return &in.ID
		}(),
	}

	for k, v := range filters {
		if v != nil {
			sb = sb.Where(sb.Equal(k, v))
		}
	}
	return sb
}

func (r *repository) GetBookList(ctx context.Context, in domain.GetBookListRequest) (*domain.GetBookListResponse, error) {
	ctx = r.queryUtil.SetPoolTx(ctx)

	sb := r.sqlbuilder.NewSelectBuilder().Select(`*`).From(constant.BookTable)

	sb = setWhere(in, sb)

	var pageOffsetRequest basemodel.PaginationOffsetRequest
	if in.PaginationRequest != nil {
		pageOffsetRequest = basemodel.PaginationOffsetRequest{
			PageId:   in.PaginationRequest.PageId,
			PageSize: in.PaginationRequest.PageSize,
		}
	}
	countSql, countArgs := sb.Build()
	sb = buildPagination(sb, pageOffsetRequest)

	query, args := sb.Build()
	bookEnties := make([]entity.Book, 0)
	if err := r.queryUtil.Query(ctx, &bookEnties, query, args); err != nil {
		return nil, err
	}
	books := typeconvertutil.ProcessItem(bookEnties, mapper.BookEntityToDomain)

	countQuery := buildCountQuery(countSql)
	totalElements, err := r.count(ctx, countQuery, countArgs)
	if err != nil {
		return nil, err
	}

	paginationResp := setPagination(pageOffsetRequest, *totalElements)
	return &domain.GetBookListResponse{
		Books:              books,
		PaginationResponse: paginationResp,
	}, nil
}
