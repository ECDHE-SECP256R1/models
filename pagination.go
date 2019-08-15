package models

type PaginationInfo struct {
	TotalCount uint64 `json:"total_count,omitempty"`
	PageCount  uint64 `json:"page_count,omitempty"`
	Page       uint64 `json:"page,omitempty"`
	PerPage    uint64 `json:"per_page,omitempty"`
	HasNext    bool   `json:"has_next,omitempty"`
}

type PaginateResult struct {
	Items      []IBaseModel   `json:"items,omitempty"`
	Pagination PaginationInfo `json:"pagination,omitempty"`
}