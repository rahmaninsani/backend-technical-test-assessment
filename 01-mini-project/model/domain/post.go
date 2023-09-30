package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;onUpdate:CASCADE;onDelete:CASCADE"`
	Title     string    `gorm:"type:varchar(200);not null"`
	Content   string    `gorm:"type:text;not null"`
	Slug      string    `gorm:"type:varchar(200);not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	PostCategories []PostCategory
	PostTags       []PostTag
}
