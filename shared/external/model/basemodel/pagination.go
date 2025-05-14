package basemodel

type PaginationOffsetRequest struct {
	PageId   int `json:"page_id"`
	PageSize int `json:"page_size"`
}

type PaginationOffsetResponse struct {
	PageId        int `json:"page_id"`
	PageSize      int `json:"page_size"`
	LastPage      int `json:"last_page"`
	TotalElements int `json:"total_elements"`
}
