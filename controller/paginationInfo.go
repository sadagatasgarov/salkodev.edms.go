package controller

// Pagination info block used in response
type PaginationInfo struct {
	TotalRecords int64 `json:"total_records"`
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	NextPage     int   `json:"next_page"`
	PrevPage     int   `json:"prev_page"`
}
