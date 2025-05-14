package repository

import (
	"fmt"
	"math"
	"yim-go/shared/external/model/basemodel"

	"github.com/huandu/go-sqlbuilder"
)

func buildPagination(sb *sqlbuilder.SelectBuilder, pg basemodel.PaginationOffsetRequest) *sqlbuilder.SelectBuilder {
	if pg.PageId < 1 || pg.PageSize == 0 {
		return sb
	}

	if pg.PageSize != -1 {
		sb = sb.Limit(pg.PageSize)
	}

	if pg.PageId-1 > 0 {
		sb.Offset((pg.PageId - 1) * pg.PageSize)
	}
	return sb
}

func calculateLastPage(totalRow, pageSize int) int {
	return int(math.Ceil((float64(totalRow) / float64(pageSize))))
}

func setPagination(in basemodel.PaginationOffsetRequest, totalElements int) basemodel.PaginationOffsetResponse {
	// Case not found on first page
	if totalElements == 0 {
		return basemodel.PaginationOffsetResponse{
			PageId:        1,
			PageSize:      in.PageSize,
			LastPage:      1,
			TotalElements: 0,
		}
	}

	// Case page size = -1 all some value for get all
	if in.PageSize < 0 || in.PageId < 1 {
		in.PageId = 1
		in.PageSize = totalElements
	}

	lastPage := calculateLastPage(totalElements, in.PageSize)
	if in.PageId <= 0 {
		lastPage = 1
	}

	return basemodel.PaginationOffsetResponse{
		PageId:        in.PageId,
		PageSize:      in.PageSize,
		LastPage:      lastPage,
		TotalElements: totalElements,
	}
}

func buildCountQuery(sql string) string {
	return fmt.Sprintf(`SELECT COUNT(*) FROM (%s) AS total_count`, sql)
}
