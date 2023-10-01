package web

import "github.com/google/uuid"

type PostCreateRequest struct {
	CategoryId uuid.UUID `json:"category_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
}
