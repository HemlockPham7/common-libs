package dto

type SuccessResponse[Data any] struct {
	Message    string      `json:"message,omitempty"`
	Data       Data        `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Page  int   `json:"page,omitempty" example:"1"`
	Limit int   `json:"limit,omitempty" example:"10"`
	Total int64 `json:"total,omitempty" example:"100"`
}
