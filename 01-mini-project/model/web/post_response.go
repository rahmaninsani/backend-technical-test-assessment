package web

import (
	"github.com/google/uuid"
	"time"
)

type PostAuthorResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type PostCategoryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type PostResponse struct {
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Slug      string               `json:"slug"`
	Category  PostCategoryResponse `json:"category"`
	Tags      []string             `json:"tags"`
	Author    PostAuthorResponse   `json:"author"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
