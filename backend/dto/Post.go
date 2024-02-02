package dto

type Post struct {
	Content string `json:"content" validate:"required"`
	Attach  string `json:"attach"`
}
