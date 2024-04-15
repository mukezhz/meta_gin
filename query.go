package meta_gin

type PaginationRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type PaginatedResult[T any] struct {
	Total   int64 `json:"total"`
	Items   []T   `json:"items"`
	HasNext bool  `json:"has_next"`
}
