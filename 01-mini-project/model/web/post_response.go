package web

import (
	"time"
)

type PostAuthorResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type PostResponse struct {
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Author    PostAuthorResponse `json:"author"`
	Slug      string             `json:"slug"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
