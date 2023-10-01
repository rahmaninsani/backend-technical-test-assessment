package web

import "github.com/google/uuid"

type CategoryCreateRequest struct {
	Name string `json:"name"`
}

type CategoryFindRequest struct {
	Id uuid.UUID `json:"category_id"`
}
