package web

import "github.com/google/uuid"

type PostCreateRequest struct {
	CategoryId uuid.UUID `json:"category_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Tags       []string  `json:"tags"`
}

type PostUpdateRequest struct {
	CategoryId uuid.UUID `json:"category_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Slug       string    `json:"slug"`
	Tags       []string  `json:"tags"`
}
